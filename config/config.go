package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config : main config struct
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	HTTPHost              string
	HTTPPort              string
	HTTPReadHeaderTimeout time.Duration
	HTTPWriteTimeout      time.Duration
	HTTPIdleTimeout       time.Duration

	AppName      string
	DebugMode    bool
	SQLDebugMode bool
}

// HTTPAddr : get full address
func (c Config) HTTPAddr() string {
	return (c.HTTPHost + ":" + c.HTTPPort)
}

// LoadConfig : load config from Config.toml
func LoadConfig() (*Config, error) {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     getEnv("CHL_DBHOST", viper.GetString("db.host")),
		DBPort:     getEnv("CHL_DBPORT", viper.GetString("db.port")),
		DBUser:     getEnv("CHL_DBUSER", viper.GetString("db.user")),
		DBPassword: getEnv("CHL_DBPASS", viper.GetString("db.password")),
		DBName:     getEnv("CHL_DBNAME", viper.GetString("db.dbname")),

		HTTPHost:              viper.GetString("server.http-host"),
		HTTPPort:              viper.GetString("server.http-port"),
		HTTPReadHeaderTimeout: viper.GetDuration("server.read-header-timeout") * time.Second,
		HTTPWriteTimeout:      viper.GetDuration("server.write-timeout") * time.Second,
		HTTPIdleTimeout:       viper.GetDuration("server.idle-timeout") * time.Second,

		AppName:      viper.Get("app-name").(string),
		DebugMode:    viper.Get("log.debug-mode").(bool),
		SQLDebugMode: viper.Get("log.sql-debug-mode").(bool),
	}, nil
}

// getEnv : Get env variable when exists
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
