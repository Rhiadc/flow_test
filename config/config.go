package config

import (
	"os"
)

type Config struct {
	Topic   string
	Brokers string
	Group   string
}

func New() *Config {

	config := &Config{
		Topic:   os.Getenv("KAFKA_TOPIC"),
		Brokers: os.Getenv("KAFKA_BROKERS"),
		Group:   os.Getenv("KAFKA_GROUP"),
	}
	return config
}
