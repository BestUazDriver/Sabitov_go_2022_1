package config

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Name string
	URI  string
}

func MongoSetUp(ctx context.Context, cancel context.CancelFunc) (*mongo.Database, error) {
	config := &MongoConfig{}
	err := viper.UnmarshalKey("mongo.database", config)
	if err != nil {
		return nil, err
	}
	client, errClient := mongo.NewClient(options.Client().ApplyURI(config.URI))
	if errClient != nil {
		return nil, errClient
	}
	errClientConnect := client.Connect(ctx)
	if errClientConnect != nil {
		return nil, errClientConnect
	}
	errPing := client.Ping(context.Background(), nil)
	if errPing != nil {
		return nil, errPing
	}

	defer cancel()
	return client.Database(config.Name), nil
}
