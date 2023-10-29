package util

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// Config Stores all configuration of the application.
// The values are nread by viper from a config file or an environment variable
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
	MINIO_ACCESS_KEY      string        `mapstructure:"MINIO_ACCESS_KEY"`
	MINIO_SECRET_KEY      string        `mapstructure:"MINIO_SECRET_KEY"`
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
	viper.SetConfigName(".env")
	viper.SetConfigType("env") // can also use json or xml

	// Override values from app.env file with environment variables if they exist
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	var config Config
	err = viper.Unmarshal(&config)

	return config
}
