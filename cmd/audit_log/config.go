package main

import (
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type appConfig struct {
	Debug       bool   `mapstructure:"debug"`
	ServiceName string `mapstructure:"service-name"`

	StartupTimeout  time.Duration `mapstructure:"startup-timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown-timeout"`
}

type mongoConfig struct {
	DSN                 string `mapstructure:"mongo-dsn"`
	Database            string `mapstructure:"mongo-db"`
	AuditLogsCollection string `mapstructure:"mongo-audit-logs-collection"`
}

type Config struct {
	App   appConfig   `mapstructure:",squash"`
	Mongo mongoConfig `mapstructure:",squash"`
}

func ReadConfig() (*Config, error) {
	_ = godotenv.Load() // nolint

	pflag.Bool("debug", false, "Enable debug")
	pflag.String("service-name", "auditlog", "Service name")
	pflag.Duration("startup-timeout", 10*time.Second, "Timeout until application should be started")
	pflag.Duration("shutdown-timeout", 15*time.Second, "Timeout until application should be stopped")

	pflag.String("mongo-dsn", "mongodb://127.0.0.1:27017", "Mongo DSN")
	pflag.String("mongo-db", "auditlog", "Mongo database for audit logs") // TODO rename it
	pflag.String("mongo-audit-logs-collection", "audit_logs", "Mongo collection name for players")

	pflag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return nil, err
	}

	var config Config

	err = viper.Unmarshal(&config)

	return &config, err
}
