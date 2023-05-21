package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type DBConnection struct {
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewDBConnection(uri string) (*DBConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &DBConnection{
		Client: client,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

func (conn *DBConnection) Ping() error {
	if err := conn.Client.Ping(conn.Ctx, readpref.Primary()); err != nil {
		return err
	}
	logrus.Print("Connected successfully")

	return nil
}

func (conn *DBConnection) Close() {
	defer conn.Cancel()
	defer func() {
		if err := conn.Client.Disconnect(conn.Ctx); err != nil {
			panic(err)
		}
	}()
}

func (conn *DBConnection) Query(dataBase, col string, query, field interface{}) (*mongo.Cursor, error) {
	collection := conn.Client.Database(dataBase).Collection(col)
	result, err := collection.Find(conn.Ctx, query, options.Find().SetProjection(field))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (conn *DBConnection) InsertOne(dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := conn.Client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(conn.Ctx, doc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (conn *DBConnection) UpdateOne(dataBase, col string, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := conn.Client.Database(dataBase).Collection(col)
	result, err := collection.UpdateOne(conn.Ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (conn *DBConnection) CancelDB() {
	conn.Cancel()
	conn.Client.Disconnect(conn.Ctx)
}
