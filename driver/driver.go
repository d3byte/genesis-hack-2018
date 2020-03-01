package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	HostKey     = string("hostKey")
	UsernameKey = string("usernameKey")
	PasswordKey = string("passwordKey")
	DatabaseKey = string("databaseKey")
)

type DB struct {
	Mgo *mongo.Database
}

var dbConn = &DB{}

// Connect to MongoDB ...
func Connect(ctx context.Context) (*DB, error) {
	dbSource := fmt.Sprintf(`mongodb://%s:%s@%s`,
		ctx.Value(UsernameKey).(string),
		ctx.Value(PasswordKey).(string),
		ctx.Value(HostKey).(string),
	)

	d, err := mongo.NewClient(options.Client().ApplyURI(dbSource))
	if err != nil {
		panic(err)
	}

	err = d.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client couldn't connect with background context: %v", err)
	}

	dbConn.Mgo = d.Database(ctx.Value(DatabaseKey).(string))

	return dbConn, err
}
