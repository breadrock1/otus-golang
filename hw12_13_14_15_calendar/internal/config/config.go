package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Logger   LoggerConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type LoggerConfig struct {
	Level         string
	FilePath      string
	EnableFileLog bool
}

type ServerConfig struct {
	Host     string
	HostPort int
	GrpcPort int
}

type DatabaseConfig struct {
	EnableInMemory bool
	PostgresURL    string
}

func NewConfig(filePath string) (*Config, error) {
	config := &Config{}

	viperInstance := viper.New()
	viperInstance.AutomaticEnv()
	viperInstance.SetConfigFile(filePath)

	viperInstance.SetDefault("logger.Level", "INFO")
	viperInstance.SetDefault("logger.FilePath", "./logs/app.log")

	viperInstance.SetDefault("server.Host", "0.0.0.0")
	viperInstance.SetDefault("server.HostPort", 2895)
	viperInstance.SetDefault("server.GrpcPort", 3895)

	viperInstance.SetDefault("database.EnableInMemory", true)

	if err := viperInstance.ReadInConfig(); err != nil {
		confErr := fmt.Errorf("failed while reading config file %s: %s", filePath, err)
		return config, confErr
	}

	if err := viperInstance.Unmarshal(config); err != nil {
		confErr := fmt.Errorf("failed while unmarshaling config file %s: %s", filePath, err)
		return config, confErr
	}

	return config, nil
}
