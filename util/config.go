package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config Stores all configuration of the application.
// The values are nread by viper from a config file or an environment variable
type Config struct {
	ENV                 string        `mapstructure:"ENV"` // prod, local
	ALLOW_ORIGIN        string        `mapstructure:"ALLOW_ORIGIN"`
	ALLOW_ORIGIN_LAN    string        `mapstructure:"ALLOW_ORIGIN_LAN"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	MINIO_ACCESS_KEY    string        `mapstructure:"MINIO_ACCESS_KEY"`
	MINIO_SECRET_KEY    string        `mapstructure:"MINIO_SECRET_KEY"`
	MINIO_ENDPOINT      string        `mapstructure:"MINIO_ENDPOINT"`
}

func IsProdEnv(conf Config) bool {
	return conf.ENV == "prod"
}

func IsLocalEnv(conf Config) bool {
	return conf.ENV == "local"
}

// LoadConfig reads configuration from file or environment variable
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env") // can also use json or xml

	// Override values from app.env file with environment variables if they exist
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
