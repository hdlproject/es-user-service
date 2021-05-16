package database

import (
	"context"
	"fmt"
	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	MongoClient struct {
		Client  *mongo.Client
		Context context.Context
		DB      *mongo.Database
	}
)

const (
	mongoUriTemplate = "mongodb://%s:%s"
)

var mongoClient *MongoClient

func GetMongoDB(mongoConfig config.EventStorage) (*MongoClient, error) {
	if mongoClient == nil {
		client, err := newMongoClient(mongoConfig.Host,
			mongoConfig.Port,
			mongoConfig.Username,
			mongoConfig.Password,
			mongoConfig.Name,
		)
		if err != nil {
			return nil, helper.WrapError(err)
		}

		mongoClient = client
	}

	return mongoClient, nil
}

func newMongoClient(host, port, username, password, dbName string) (*MongoClient, error) {
	credential := options.Credential{
		Username: username,
		Password: password,
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx,
		options.Client().
			ApplyURI(fmt.Sprintf(mongoUriTemplate, host, port)).
			SetAuth(credential))
	if err != nil {
		return nil, helper.WrapError(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return &MongoClient{
		Client:  client,
		Context: ctx,
		DB:      client.Database(dbName),
	}, nil
}
