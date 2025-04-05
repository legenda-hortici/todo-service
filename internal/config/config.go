package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Database   `yaml:"database"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Name     string `yaml:"name" env-required:"true"`
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH env variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read config:", err)
	}

	return &cfg
}



