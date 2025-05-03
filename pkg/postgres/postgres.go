// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"local/order-service/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"

	defaults "github.com/creasty/defaults"

	"github.com/sirupsen/logrus"
)

type pgxLogger struct {
	lg logger.Interface
}

func (l pgxLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	switch level {
	case tracelog.LogLevelTrace:
		l.lg.Trace(msg, data)
	case tracelog.LogLevelDebug:
		l.lg.Debug(msg, data)
	case tracelog.LogLevelInfo:
		l.lg.Info(msg, data)
	case tracelog.LogLevelWarn:
		l.lg.Warn(msg, data)
	case tracelog.LogLevelError:
		l.lg.Error(msg, data)
	case tracelog.LogLevelNone:
		return
	}
}

type Postgres struct {
	Pool    *pgxpool.Pool
	poolCfg *pgxpool.Config
	cfg     Config
}

func Connect(cfg Config, logger logger.Interface) (*Postgres, error) {

	if logger == nil {
		return nil, fmt.Errorf("logger not initialized")
	}

	lg := pgxLogger{lg: logger}

	pg := &Postgres{
		cfg: cfg,
	}

	defaults.Set(&cfg)

	config, err := pgxpool.ParseConfig(cfg.PostgreDSN)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse config: %v", err)
	}

	config.MaxConnLifetime = cfg.MaxConnLifetime
	config.MaxConnIdleTime = cfg.MaxConnIdleTime
	config.MaxConns = cfg.MaxConns
	config.MinConns = cfg.MinConns
	config.ConnConfig.ConnectTimeout = cfg.ConnectTimeout

	logLevel, err := tracelog.LogLevelFromString(cfg.LogLevel)

	if err != nil {
		logrus.Errorf("incorrect log level %s. fallback to debug", cfg.LogLevel)
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   lg,
		LogLevel: logLevel,
	}

	pg.poolCfg = config

	ctx, cancel := context.WithTimeout(context.Background(), config.ConnConfig.ConnectTimeout)
	defer cancel()

	for cfg.ConnectAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(ctx, config)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", cfg.ConnectAttempts)

		time.Sleep(cfg.ReconnectTimeout)

		cfg.ConnectAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	// Check connection
	if err := pg.Pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
