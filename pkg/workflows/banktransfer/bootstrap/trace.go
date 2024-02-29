package bootstrap

import (
	"context"
	"fmt"

	"github.com/anhgeeky/go-temporal-labs/core/config"
	"github.com/anhgeeky/go-temporal-labs/core/trace"
	"github.com/anhgeeky/go-temporal-labs/core/trace/otel"
)

func ConfigTracing(srvCfg *ServerConfig, cfg config.Configure) {
	endpoint := cfg.GetString("TRACE_ENDPOINT")
	otel.SetGlobalTracer(context.Background(), srvCfg.Name, endpoint)
}

func GetTracer(srvCfg *ServerConfig, cfg config.Configure) trace.Tracer {
	endpoint := cfg.GetString("TRACE_ENDPOINT")
	tracer, err := otel.NewOpenTelemetryTracer(context.Background(), srvCfg.Name, endpoint)
	if err != nil {
		panic(fmt.Errorf("Failed to create tracer object: %w", err))
	}
	return tracer
}
