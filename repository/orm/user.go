package orm

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google-service/models"
	"google-service/repository"
	"time"
)

type MongoUserRepo struct {
	Conn *mongo.Database
}

func (m MongoUserRepo) NewSQLRepo() repository.UserRepo {
	return &m
}

func (m MongoUserRepo) GenerateToken(ctx context.Context) (string, error) {
	c := &models.User{ID: primitive.NewObjectID()}
	err := m.createToken(c)
	if err != nil {
		return "", err
	}

	return c.Token, nil
}

func (m *MongoUserRepo) Fetch(ctx context.Context, num int64) ([]*models.UserWithHiddenFields, error) {
	res, err := m.Conn.Collection("users").Find(ctx, bson.D{}, options.Find().SetLimit(num), options.Find().SetSort(bson.M{"time": -1}))
	if err != nil {
		return nil, fmt.Errorf("fetch data: couldn't list all data: %v", err)
	}
	defer res.Close(ctx)

	payload := make([]*models.UserWithHiddenFields, 0)

	for res.Next(ctx) {
		data := new(models.UserWithHiddenFields)

		if err = res.Decode(&data); err != nil {
			return nil, fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
		}

		data.Token = ""
		payload = append(payload, data)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("fetch data: all items couldn't be listed: %v", err)
	}

	return payload, nil
}

func (m *MongoUserRepo) GetUserById(ctx context.Context, id primitive.ObjectID) (*models.UserWithHiddenFields, error) {
	var payload *models.UserWithHiddenFields

	err := m.Conn.Collection("users").FindOne(ctx, bson.D{{"_id", id}}).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
	}

	return payload, nil
}

func (m *MongoUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var payload *models.User

	err := m.Conn.Collection("users").FindOne(ctx, bson.D{{"email", email}}).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("fetch data: couldn't make item ready for display: %v", err)
	}

	err = m.createToken(payload)
	if err != nil {
		return nil, fmt.Errorf("create user: token couldn't be created: %v", err)
	}

	return payload, nil
}

func (m *MongoUserRepo) Create(ctx context.Context, c *models.User) (string, error) {
	c.ID = primitive.NewObjectID()

	err := m.createToken(c)
	if err != nil {
		return "", fmt.Errorf("create user: token couldn't be created: %v", err)
	}

	res, err := m.Conn.Collection("users").InsertOne(ctx, c)
	if err != nil {
		return "", fmt.Errorf("create user: user couldn't be created: %v", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *MongoUserRepo) createToken(c *models.User) error {
	role := ""

	if c.Role == 0 {
		role = "USER"
	} else if c.Role == 1 {
		role = "CREATOR"
	}

	claims := jwt.MapClaims{
		"id":       c.ID,
		"role":		role,
		"exp":      time.Now().UTC().Add(time.Hour).Unix(),
		"orig_iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("SESSION_SECRET"))
	if err != nil {
		fmt.Println(err)
		return err
	}

	c.Token = tokenString

	return nil
}
