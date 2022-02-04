package http

import (
	"net/http"

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
