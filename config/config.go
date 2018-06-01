package config

import (
	"fmt"
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

	AppName   string
	DebugMode bool
}

// HTTPAddr : get full address
func (c Config) HTTPAddr() string {
	return ("http://" + c.HTTPHost + ":" + c.HTTPPort)
}

// LoadConfig : load config from config.toml
func LoadConfig() (*Config, error) {
	c := viper.New()
	c.SetConfigName("config")
	c.AddConfigPath("./config")

	err := c.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     getEnv("CHL_DBHOST", c.GetString("db.host")),
		DBPort:     getEnv("CHL_DBPORT", c.GetString("db.port")),
		DBUser:     getEnv("CHL_DBUSER", c.GetString("db.user")),
		DBPassword: getEnv("CHL_DBPASS", c.GetString("db.password")),
		DBName:     getEnv("CHL_DBNAME", c.GetString("db.dbname")),

		HTTPHost:              c.GetString("server.http-host"),
		HTTPPort:              c.GetString("server.http-port"),
		HTTPReadHeaderTimeout: c.GetDuration("server.read-header-timeout") * time.Second,
		HTTPWriteTimeout:      c.GetDuration("server.write-timeout") * time.Second,
		HTTPIdleTimeout:       c.GetDuration("server.idle-timeout") * time.Second,

		AppName:   c.Get("app-name").(string),
		DebugMode: c.Get("log.debug-mode").(bool),
	}, nil
}

// MustLoadConfig : load config from config.toml, panics on error
func MustLoadConfig() *Config {
	viper, err := LoadConfig()
	if err != nil {
		panic("cant load config file:  " + err.Error())
	}
	return viper

}

// GetConnString : get database connection string
func (c *Config) GetDBConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

// getEnv : Get env variable when exists
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
