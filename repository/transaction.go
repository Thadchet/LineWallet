package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	db *mongo.Database
}

type ITransactionRepo interface {
	Insert() error
	InsertTransaction(m interface{}) error
}

func (t TransactionRepo) Insert() error {

	m := bson.M{"a": "b"}
	_, err := t.db.Collection("transactions").InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionRepo) InsertTransaction(m interface{}) error {
	_, err := t.db.Collection("transactions").InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}
	return nil
}
