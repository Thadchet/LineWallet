package repository

import (
	"context"
	"fmt"
	"line-wallet/models"
	"line-wallet/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	db *mongo.Database
}

type ITransactionRepo interface {
	Insert() error
	InsertTransaction(m interface{}) error
	GetTransactions(line_user_id string) ([]models.Transaction, error)
	GetTransactionByID(ID string) (*models.Transaction, error)
	UpdateTransactionByID(amount, category, memo string, id string) (*mongo.UpdateResult, error)
	FilterTransactionCurrentMonth() ([]models.Transaction, error)
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

func (t TransactionRepo) GetTransactionByID(ID string) (*models.Transaction, error) {
	var result *models.Transaction
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": objID,
	}
	if err := t.db.Collection("transactions").FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t TransactionRepo) UpdateTransactionByID(amount, category, memo, id string) (*mongo.UpdateResult, error) {

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objID,
	}
	update := bson.M{
		"$set": bson.M{
			"amount":   amount,
			"category": category,
			"memo":     memo,
		},
	}
	updated, err := t.db.Collection("transactions").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (t TransactionRepo) FilterTransactionCurrentMonth() ([]models.Transaction, error) {
	currentMonth := time.Now().Month().String()
	amountDay := utils.GetMaxDay(currentMonth)
	dateFrom := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	dateTo := time.Date(time.Now().Year(), time.Now().Month(), amountDay, 0, 0, 0, 0, time.Local)

	fmt.Println(dateFrom)
	fmt.Println(dateTo)

	filter := bson.M{
		"createdat": bson.M{
			"$gte": dateFrom,
			"$lte": dateTo,
		},
		"type": "txn",
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
