package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type DBConnection struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func NewDBConnection(uri string) (*DBConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &DBConnection{
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (conn *DBConnection) Ping() error {
	if err := conn.client.Ping(conn.ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Print("Connected successfully")

	return nil
}

func (conn *DBConnection) Close() {
	defer conn.cancel()
	defer func() {
		if err := conn.client.Disconnect(conn.ctx); err != nil {
			panic(err)
		}
	}()
}

func (conn *DBConnection) Query(dataBase, col string, query, field interface{}) (*mongo.Cursor, error) {
	collection := conn.client.Database(dataBase).Collection(col)
	result, err := collection.Find(conn.ctx, query, options.Find().SetProjection(field))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (conn *DBConnection) InsertOne(dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := conn.client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(conn.ctx, doc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (conn *DBConnection) UpdateOne(dataBase, col string, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := conn.client.Database(dataBase).Collection(col)
	result, err := collection.UpdateOne(conn.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (conn *DBConnection) Cancel() {
	conn.cancel()
	conn.client.Disconnect(conn.ctx)
}
