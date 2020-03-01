package orm

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google-service/models"
)

type MongoMiptRepo struct {
	Conn *mongo.Database
}

func (m MongoMiptRepo) NewSQLRepo() *MongoMiptRepo {
	return &m
}

func (m MongoMiptRepo) Fetch(ctx context.Context, num int64) (*models.ConfigInterface, error) {
	panic("implement me")
}

func (m MongoMiptRepo) GetMiptById(ctx context.Context, id primitive.ObjectID) (*models.ConfigInterface, error) {
	var payload *models.ConfigInterface

	err := m.Conn.Collection("configs").FindOne(ctx, bson.D{{"_id", id}}).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
	}

	return payload, nil
}

func (m MongoMiptRepo) Create(ctx context.Context, s *models.ConfigInterface) (string, error) {
	res, err := m.Conn.Collection("configs").InsertOne(ctx, s)
	if err != nil {
		return "", fmt.Errorf("create config: config couldn't be created: %v", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
