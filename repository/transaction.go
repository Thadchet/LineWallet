package repository

import (
	"context"
	"fmt"
	"line-wallet/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	db *mongo.Database
}

type ITransactionRepo interface {
	Insert() error
	InsertTransaction(m interface{}) error
	GetTransactions(line_user_id string) ([]models.Transaction, error)
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

func (t TransactionRepo) GetTransactions(line_user_id string) ([]models.Transaction, error) {

	filter := bson.M{
		"lineuserid": line_user_id,
	}
	cursors, err := t.db.Collection("transactions").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var result []models.Transaction
	if err = cursors.All(context.TODO(), &result); err != nil {
		fmt.Println(err.Error())
	}
	return result, nil
}
