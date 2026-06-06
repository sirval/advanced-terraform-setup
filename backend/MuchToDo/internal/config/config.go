package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	ServerPort         string `mapstructure:"PORT"`
	MongoURI          string `mapstructure:"MONGO_URI"`
	DBName            string `mapstructure:"DB_NAME"`
	JWTSecretKey      string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int    `mapstructure:"JWT_EXPIRATION_HOURS"`

	EnableCache   bool   `mapstructure:"ENABLE_CACHE"`
	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	LogLevel  string `mapstructure:"LOG_LEVEL"`
	LogFormat string `mapstructure:"LOG_FORMAT"`

	Integration bool `mapstructure:"INTEGRATION"`

	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
	CookieDomains  []string `mapstructure:"COOKIE_DOMAINS"`
	SecureCookie   bool     `mapstructure:"SECURE_COOKIE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	envKeys := []string{
		"PORT",
		"MONGO_URI",
		"DB_NAME",
		"JWT_SECRET_KEY",
		"JWT_EXPIRATION_HOURS",
		"ENABLE_CACHE",
		"REDIS_ADDR",
		"REDIS_PASSWORD",
		"LOG_LEVEL",
		"LOG_FORMAT",
		"INTEGRATION",
		"ALLOWED_ORIGINS",
		"COOKIE_DOMAINS",
		"SECURE_COOKIE",
	}

	for _, key := range envKeys {
		if bindErr := viper.BindEnv(key); bindErr != nil {
			return config, bindErr
		}
	}

	// Default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DB_NAME", "much_todo_db")
	viper.SetDefault("ENABLE_CACHE", false)
	viper.SetDefault("JWT_EXPIRATION_HOURS", 72)
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("LOG_FORMAT", "json")
	viper.SetDefault("INTEGRATION", true)
	viper.SetDefault("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:5173,http://localhost:5174")
	viper.SetDefault("COOKIE_DOMAINS", "localhost,127.0.0.1")
	viper.SetDefault("SECURE_COOKIE", false)

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	// Convert comma-separated env values into slices
	config.AllowedOrigins = splitCommaSeparated(viper.GetString("ALLOWED_ORIGINS"))
	config.CookieDomains = splitCommaSeparated(viper.GetString("COOKIE_DOMAINS"))

	return config, nil
}

func splitCommaSeparated(value string) []string {
	if value == "" {
		return []string{}
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}