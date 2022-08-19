package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/igilgyrg/crypto/pkg/logging"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type Config struct {
	ListenType      string `env:"LISTEN_TYPE" envDefault:"port"`
	BindIp          string `env:"BIND_IP" envDefault:"0.0.0.0"`
	Port            string `env:"PORT" envDefault:"80"`
	MongoHost       string `env:"MONGO_HOST"`
	MongoPort       string `env:"MONGO_PORT"`
	MongoDatabase   string `env:"MONGO_DATABASE"`
	MongoUsername   string `env:"MONGO_USERNAME"`
	MongoPassword   string `env:"MONGO_PASSWORD"`
	MongoCollection string `env:"MONGO_COLLECTION"`
	FtxApiKey       string `env:"FTX_API_KEY"`
	FtxApiSecret    string `env:"FTX_API_SECRET"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("load configurations")
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
