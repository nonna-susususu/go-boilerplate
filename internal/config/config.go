package config

import "github.com/fastworkco/common-go/telemetry/v1"

type Config struct {
	Telemetry        telemetry.Config
	InternalServices InternalServices
	AppConfig        AppConfig
	RemoteConfig     RemoteConfig
	DatabaseConfig   DatabaseConfig
}

type AppConfig struct {
	AppName    string   `env:"APP_NAME,required"`
	Env        string   `env:"ENV,required"`
	Port       int      `env:"PORT" envDefault:"3000"`
	Cors       []string `env:"APP_CORS_ALLOW_ORIGIN,required"`
	CorsHeader []string `env:"APP_CORS_ALLOW_HEADER"`
}

type RemoteConfig struct {
	FastworkAuthProviderURL string `env:"FASTWORK_AUTH_URL"`
}

type DatabaseConfig struct {
	DBHost     string `env:"DB_HOST" envDefault:"postgres"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:""`
	DBName     string `env:"DB_NAME" envDefault:"go-boilerplate"`
}

type InternalServices struct{}
