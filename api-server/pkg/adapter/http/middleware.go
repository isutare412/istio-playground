package http

import (
	"errors"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	log "github.com/sirupsen/logrus"
)

type responseLogger struct {
	http.ResponseWriter
	status int
	length int
}

func (l *responseLogger) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

func (l *responseLogger) Write(b []byte) (int, error) {
	l.length += len(b)
	return l.ResponseWriter.Write(b)
}

func structAccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := responseLogger{ResponseWriter: w, status: http.StatusOK}
		h.ServeHTTP(&logger, r)

		log.WithFields(log.Fields{
			"addr":   r.RemoteAddr,
			"method": r.Method,
			"url":    r.URL.String(),
			"status": logger.status,
			"length": logger.length,
		}).Info("access")
	})
}

func plainAccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := responseLogger{ResponseWriter: w, status: http.StatusOK}
		h.ServeHTTP(&logger, r)

		log.Infof("%s - \"%s %s\" %d %d",
			r.RemoteAddr, r.Method, r.URL.String(), logger.status, logger.length)
	})
}

func tracing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var span opentracing.Span
		wireCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		if err != nil {
			if !errors.Is(err, opentracing.ErrSpanContextNotFound) {
				log.Errorf("invalid span detected: %v", err)
				responseError(w, http.StatusInternalServerError, "invalid span detected")
				return
			}

			span = opentracing.GlobalTracer().StartSpan("http.middleware.tracing")
			err = span.Tracer().Inject(
				span.Context(),
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(r.Header),
			)
			if err != nil {
				log.Errorf("failed to inject new span: %v", err)
				responseError(w, http.StatusInternalServerError, "failed to inject new span")
				return
			}
		} else {
			span = opentracing.GlobalTracer().StartSpan(
				"http.middleware.tracing",
				ext.RPCServerOption(wireCtx),
			)
		}
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(r.Context(), span)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
