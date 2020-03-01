package orm

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google-service/models"
	"google.golang.org/api/sheets/v4"
)

type MongoStateRepo struct {
	Conn *mongo.Database
}

func (m MongoStateRepo) NewSQLRepo() *MongoStateRepo {
	return &m
}

func (m MongoStateRepo) GetStateById(ctx context.Context, s *models.StateInterface) (*models.StateInterface, error) {
	var payload *models.StateInterface

	err := m.Conn.Collection("states").FindOne(ctx, bson.D{{"userid", s.UserID}, {"configid", s.ConfigID}}).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
	}

	return payload, nil
}

func (m MongoStateRepo) Create(ctx context.Context, s *models.StateInterface) (string, error) {
	res, err := m.Conn.Collection("states").InsertOne(ctx, s)
	if err != nil {
		return "", fmt.Errorf("create config: config couldn't be created: %v", err)
	}

	if s.Completed == true {
		conf, _ := m.getMiptById(ctx, s.ConfigID)
		spread := models.Spreadsheet{
			Token: conf.PublicToken,
		}

		var arr []interface{}
		arr = append(arr, s.UserID)

		for i := 0; i < len(s.Answers); i++ {
			arr = append(arr, s.Answers[i].Value)
		}
		spread.Data = append(spread.Data, arr)

		err = m.appendData(ctx, &spread)
		if err != nil {
			return "", err
		}
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m MongoStateRepo) CheckMiptById(ctx context.Context, id primitive.ObjectID) error {
	var payload *models.ConfigInterface

	err := m.Conn.Collection("configs").FindOne(ctx, bson.D{{"_id", id}}).Decode(&payload)
	if err != nil {
		return fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
	}

	return nil
}

func (m MongoStateRepo) appendData(ctx context.Context, s *models.Spreadsheet) error {
	srv := SetupGoogle()

	spreadsheetId := s.Token

	writeRange := "A2"

	var vr sheets.ValueRange

	for _, v := range s.Data {
		vr.Values = append(vr.Values, v)
	}

	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (m MongoStateRepo) getMiptById(ctx context.Context, id primitive.ObjectID) (*models.ConfigInterface, error) {
	var payload *models.ConfigInterface

	err := m.Conn.Collection("configs").FindOne(ctx, bson.D{{"_id", id}}).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
	}

	return payload, nil
}
