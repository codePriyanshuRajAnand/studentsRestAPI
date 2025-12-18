package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpAddress struct {
	Addr string
}

type ProjectConfig struct {
	Env          string `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HttpAddress  `yaml:"http_server"`
}

func MustLoad() *ProjectConfig {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "Path of Configuration files")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exists %s", configPath)
	}
	var cfg ProjectConfig
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		log.Fatalf("Unable to read config file: %s", err.Error())
	}
	return &cfg
}
