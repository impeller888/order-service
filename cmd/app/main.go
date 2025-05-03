package main

import (
	"log"

	"local/order-service/internal/app"
	"local/order-service/internal/app/config"
	"local/order-service/pkg/tracing"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig("")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	if cfg.Tracing.Enabled {
		tp, tpErr := tracing.JaegerTracerProvider()
		if tpErr != nil {
			log.Fatal(tpErr)
		}
		otel.SetTracerProvider(tp)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	}

	// Run
	app.Run(cfg)
}
