package functional

import (
	"fmt"
	"local/order-service/internal/app"
	"local/order-service/internal/app/config"
	"local/order-service/pkg/logger"
	"local/order-service/pkg/postgres"
	"log"
	"net/http"
	"testing"
)

var (
	handler http.Handler
	pg      *postgres.Postgres
)

func init() {
	cfg, err := config.NewConfig("../../config/config.yaml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	lg := logger.New(cfg.Log.Level)
	cfg.Metrics.Enabled = false
	cfg.Swagger.Enabled = false
	cfg.Tracing.Enabled = false

	// Repository
	pg, err = postgres.Connect(cfg.DB.PG, lg)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	// call flag.Parse() here if TestMain uses flags

	handler = app.MakeRouter(cfg, pg.Pool, lg).Engine
}

func TestMain(m *testing.M) {
	m.Run()
	pg.Close()
}
