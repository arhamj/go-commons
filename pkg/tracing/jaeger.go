package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

type Config struct {
	ServiceName string `mapstructure:"service_name"`
	HostPort    string `mapstructure:"host_port"`
	Enable      bool   `mapstructure:"enable"`
	LogSpans    bool   `mapstructure:"log_spans"`
}

func NewJaegerTracer(jaegerConfig *Config) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: jaegerConfig.ServiceName,

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &config.ReporterConfig{
			LogSpans:           jaegerConfig.LogSpans,
			LocalAgentHostPort: jaegerConfig.HostPort,
		},
	}

	return cfg.NewTracer(config.Logger(jaeger.StdLogger))
}
