package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port int `env:"PORT"`
	Host string `env:"HOST"`
	TotalServers int `env:"TOTAL_SERVERS"`

	
}

func InitConfig() (Config, error)  {
	configurationPath := os.Getenv("CONFIG_PATH")
	if configurationPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configurationPath = *flags

		if configurationPath == "" {
            log.Fatal("Error: configuration path is not provided")
        }
	}

	_, err := os.Stat(configurationPath); if os.IsNotExist(err) {
		log.Fatalf("Error: configuration file at %s not exist!", configurationPath)
	}

	var config Config

	err = cleanenv.ReadConfig(configurationPath, config); if err != nil {
		log.Fatal("Error: Failed to read configuration file")
	}

	cleanenv.ReadEnv(config)

	return config, nil
}

