package config

import (
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppCfg AppConfig
	DBCfg  DBConfig
}

type AppConfig struct {
	AppName         string
	AppHeaderName   string
	AppVersion      string
	AppPort         string
	AllowOrigins    string
	AllowMethods    string
	AllowHeaders    string
	AppReadTimeout  time.Duration
	AppWriteTimeout time.Duration
	AppIdleTimeout  time.Duration
}

type DBConfig struct {
	DBDriver string
	DBPath   string
}

func Load() *Config {
	return &Config{
		AppCfg: AppConfig{
			AppName:         getEnv("APP_NAME", "MyApp"),
			AppHeaderName:   getEnv("APP_HEADER_NAME", "X-MyApp"),
			AppVersion:      getEnv("APP_VERSION", "1.0.0"),
			AppPort:         getEnv("APP_PORT", "8080"),
			AppReadTimeout:  getEnvAsDuration("APP_READ_TIMEOUT", 10*time.Second),
			AppWriteTimeout: getEnvAsDuration("APP_WRITE_TIMEOUT", 10*time.Second),
			AppIdleTimeout:  getEnvAsDuration("APP_IDLE_TIMEOUT", 60*time.Second),
		},
		DBCfg: DBConfig{
			DBDriver: getEnv("DB_DRIVER", "sqlite3"),
			DBPath:   getEnv("GOOSE_DBSTRING", "data/myapp.db"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
