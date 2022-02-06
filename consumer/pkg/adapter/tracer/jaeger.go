package tracer

import (
	"io"

	"github.com/isutare412/istio-playground/consumer/pkg/config"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jConfig "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

func NewTracer(cfg *config.TracerConfig) (opentracing.Tracer, io.Closer, error) {
	builder := jConfig.Configuration{
		Disabled:    !cfg.Enabled,
		ServiceName: cfg.ServiceName,
		Sampler: &jConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jConfig.ReporterConfig{
			CollectorEndpoint: cfg.Endpoint,
			LogSpans:          false,
		},
	}

	b3Propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, closer, err := builder.NewTracer(
		jConfig.Logger(&jaegerLogger{}),
		jConfig.Injector(opentracing.HTTPHeaders, b3Propagator),
		jConfig.Extractor(opentracing.HTTPHeaders, b3Propagator),
		jConfig.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		return nil, nil, err
	}
	return tracer, closer, nil
}
