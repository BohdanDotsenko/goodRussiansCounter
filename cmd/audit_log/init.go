package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	connectTimeout = 10 * time.Second
)

func MustReadConfig() *Config {
	config, err := ReadConfig()
	if err != nil {
		panic(fmt.Sprintf("read config: %s", err))
	}

	return config
}

func MustCreateMongoClient(
	config *Config,
) *mongo.Client {
	opts := options.Client().SetConnectTimeout(connectTimeout)
	opts.ApplyURI(config.Mongo.DSN)

	client, err := mongo.NewClient(opts)
	if err != nil {
		panic("create mongo client")
	}

	return client
}
