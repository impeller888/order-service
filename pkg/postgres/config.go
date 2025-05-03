package postgres

import "time"

type Config struct {
	PostgreDSN           string        `yaml:"dsn" default:"host=127.0.0.1 user=cdp_task_manager password=cdp_task_manager dbname=cdp_task_manager port=5432"`
	MaxConns             int32         `default:"10"`
	MinConns             int32         `default:"2"`
	MaxConnLifetime      time.Duration `default:"30m"`
	MaxConnIdleTime      time.Duration `default:"5m"`
	ConnectTimeout       time.Duration `default:"5s"`
	ReconnectTimeout     time.Duration `default:"5s"`
	ConnectAttempts      int32         `default:"3"`
	LogLevel             string        `default:"debug"`
	PreferSimpleProtocol bool          `default:"true"`
}
