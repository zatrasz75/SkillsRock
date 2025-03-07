package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type (
	Config struct {
		Server `yaml:"httpServer"`
		PG     `yaml:"postgres"`
	}
	Server struct {
		AddrPort           string        `env:"APP_PORT" env-default:"3000"`
		AddrHost           string        `env:"APP_IP" env-default:"localhost"`
		ShutdownTime       time.Duration `env:"SHUTDOWN_TIMEOUT" env-default:"10s"`
		CORSAllowedOrigins []string      `env:"CORS_ALLOWED_ORIGINS" env-default:"localhost"`
	}

	PG struct {
		ConnStr string `env:"DB_CONNECTION_STRING"`

		User     string `yaml:"username" env:"POSTGRES_USER" env-default:"zatrasz"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgrespw"`
		Host     string `yaml:"host" env:"HOST_DB" env-default:"localhost"`
		Port     string `yaml:"port" env:"PORT_DB" env-default:"49781"`
		Name     string `yaml:"db-name" env:"POSTGRES_DB" env-default:"db_rock"`

		PoolMax      int           `yaml:"pool-max" env:"PG_POOL_MAX" env-default:"10"`
		ConnAttempts int           `yaml:"conn-attempts" env:"PG_CONN_ATTEMPTS" env-default:"5"`
		ConnTimeout  time.Duration `yaml:"conn-timeout" env:"PG_TIMEOUT" env-default:"2s"`
	}
)

func NewConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	cfg.ConnStr = initDB(cfg)

	return &cfg, nil
}

func initDB(cfg Config) string {
	if cfg.ConnStr != "" {
		return cfg.ConnStr
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}
