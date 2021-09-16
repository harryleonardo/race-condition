package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// MongoInterface ...
	MongoInterface interface {
		OpenMongoConn() *mongo.Collection
	}
)

func (d *database) OpenMongoConn() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(d.SharedConfig.MONGO.URI))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	return client
}
