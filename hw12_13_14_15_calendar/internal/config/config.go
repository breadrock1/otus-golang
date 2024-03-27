package config

type Config struct {
	Logger LoggerConfig
	// TODO
}

type LoggerConfig struct {
	Level    string
	FilePath string
}

type ServiceConfig struct {
	Host string
	Port int
}

func NewConfig() *Config {
	return &Config{}
}

// TODO
