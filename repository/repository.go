package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google-service/models"
)

type GoogleRepo interface {
	AppendData(ctx context.Context, s *models.Spreadsheet) error
	ClearData(ctx context.Context, s *models.SpreadsheetClear) error
	CreateTable(ctx context.Context) (*models.Spreadsheet, error)
	CopyTable(ctx context.Context, s *models.Spreadsheet) (*models.Spreadsheet, error)
}

type MiptRepo interface {
	Fetch(ctx context.Context, num int64) (*models.ConfigInterface, error)
	GetMiptById(ctx context.Context, id primitive.ObjectID) (*models.ConfigInterface, error)
	Create(ctx context.Context, s *models.ConfigInterface) (string, error)
}

type StateRepo interface {
	GetStateById(ctx context.Context, s *models.StateInterface) (*models.StateInterface, error)
	Create(ctx context.Context, s *models.StateInterface) (string, error)
	CheckMiptById(ctx context.Context, id primitive.ObjectID) error
}

var secretKey string

type UserRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.UserWithHiddenFields, error)
	GetUserById(ctx context.Context, id primitive.ObjectID) (*models.UserWithHiddenFields, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, c *models.User) (string, error)
	GenerateToken(ctx context.Context) (string, error)
}
