package database

import (
	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	topUpEvent struct {
		ID                 primitive.ObjectID `bson:"_id,omitempty"`
		TransactionEventID string             `bson:"transaction_event_id,omitempty"`
		UserID             uint               `bson:"user_id,omitempty"`
		Amount             uint64             `bson:"amount,omitempty"`
	}
)

type (
	topUpEventRepo struct {
		mongoClient          *MongoClient
		topUpEventCollection *mongo.Collection
	}
)

func NewTransactionEventRepo(mongoClient *MongoClient) output_port.TopUpEventRepo {
	return &topUpEventRepo{
		mongoClient:          mongoClient,
		topUpEventCollection: mongoClient.db.Collection("topup_events"),
	}
}

func (instance *topUpEventRepo) Insert(event entity.TopUpEvent) (string, error) {
	data, _ := topUpEvent{}.getData(event)
	result, err := instance.topUpEventCollection.InsertOne(mongoClient.context, data)
	if err != nil {
		return "", helper.WrapError(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (instance *topUpEventRepo) IsAlreadyApplied(event entity.TopUpEvent) (bool, error) {
	filter, _ := topUpEvent{}.getData(event)
	var data topUpEvent
	err := instance.topUpEventCollection.FindOne(mongoClient.context, bson.M{"transaction_event_id": filter.TransactionEventID}).Decode(&data)
	if err != nil {
		return false, helper.WrapError(err)
	}

	return true, nil
}

func (topUpEvent) getData(topUpEventEntity entity.TopUpEvent) (topUpEvent, error) {
	return topUpEvent{
		TransactionEventID: topUpEventEntity.TransactionEventID,
		UserID:             topUpEventEntity.UserID,
		Amount:             topUpEventEntity.Amount,
	}, nil
}
