package util

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// Config Stores all configuration of the application.
// The values are read by viper from a config file or an environment variable
type Config struct {
	ENV                   string        `mapstructure:"ENV"` // prod, local
	GIN_MODE              string        `mapstructure:"GIN_MODE"`
	ALLOW_ORIGIN          string        `mapstructure:"ALLOW_ORIGIN"`
	ALLOW_ORIGIN_LAN      string        `mapstructure:"ALLOW_ORIGIN_LAN"`
	DBDriver              string        `mapstructure:"DB_DRIVER"`
	DBSource              string        `mapstructure:"DB_SOURCE"`
	ServerAddress         string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration   time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	MINIO_ROOT_USER      string        `mapstructure:"MINIO_ROOT_USER"`
	MINIO_ROOT_PASSWORD      string        `mapstructure:"MINIO_ROOT_PASSWORD"`
	MINIO_ENDPOINT        string        `mapstructure:"MINIO_ENDPOINT"`
	MINIO_PUBLIC_ENDPOINT string        `mapstructure:"MINIO_PUBLIC_ENDPOINT"`
}

func (conf *Config) IsProdEnv() bool {
	return conf.ENV == "prod"
}

func (conf *Config) IsLocalEnv() bool {
	return conf.ENV == "local"
}

func (conf *Config) GetSameSite() http.SameSite {
	return http.SameSiteDefaultMode
	// if conf.IsLocalEnv() {
	// 	return http.SameSiteNoneMode
	// }
}

// LoadConfig reads configuration from file or environment variable
func LoadConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.AddConfigPath("/vault/secrets")

	viper.SetConfigType("env") // can also use json or xml
	viper.SetConfigName(".env")


	// Defaults
	viper.SetDefault("ENV", "prod")
	viper.SetDefault("ALLOW_ORIGIN", "https://shinypothos.com")
	viper.SetDefault("ALLOW_ORIGIN_LAN", "http://10.244.0.1")
	viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:1337")
	viper.SetDefault("ACCESS_TOKEN_DURATION", "60m")
	viper.SetDefault("GIN_MODE", "release")
	viper.SetDefault("DB_DRIVER", "postgres")


	// Override values from app.env file with environment variables if they exist
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Could not read in config - viper: ", err)
	}

	var config Config
	err = viper.Unmarshal(&config)

	return config
}
