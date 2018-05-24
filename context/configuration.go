package context

import (
	log "changelog/log"
	"os"

	"github.com/spf13/viper"
)

// Config : main config struct
type Config struct {
	AppName string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	DebugMode bool

	HTTPPort string
}

// LoadConfig : load config from Config.toml
func LoadConfig(path string) *Config {
	config := viper.New()
	config.SetConfigName("Config")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error context file: %s \n", err)
	}

	return &Config{
		AppName:    config.Get("app-name").(string),
		DBHost:     getEnv("CHL_DBHOST", config.Get("db.host").(string)),
		DBPort:     getEnv("CHL_DBPORT", config.Get("db.port").(string)),
		DBUser:     getEnv("CHL_DBUSER", config.Get("db.user").(string)),
		DBPassword: getEnv("CHL_DBPASS", config.Get("db.password").(string)),
		DBName:     getEnv("CHL_DBNAME", config.Get("db.dbname").(string)),
		DebugMode:  config.Get("log.debug-mode").(bool),
		HTTPPort:   config.Get("server.http-port").(string),
	}
}

// Get env variable when exists
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
