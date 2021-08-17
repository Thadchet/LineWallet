package repository

import (
	"context"
	"fmt"
	"line-wallet/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberRepo struct {
	db *mongo.Database
}

type IMemberRepo interface {
	CreateMember(member models.Member) error
	FindMemberByLineUserID(line_use_id string) (*models.Member, error)
	UpdateRemainingBalance(line_use_id string, amount float64) error
}

func (t MemberRepo) CreateMember(member models.Member) error {
	var temp models.Member
	options := bson.M{"lineuserid": member.LineUserID}
	if err := t.db.Collection("member").FindOne(context.TODO(), options).Decode(temp); err != nil {
		if err.Error() == "mongo: no documents in result" {
			_, err := t.db.Collection("member").InsertOne(context.TODO(), member)
			if err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("duplidate member")
	}
	return nil
}

func (t MemberRepo) FindMemberByLineUserID(line_use_id string) (*models.Member, error) {
	var result models.Member
	options := bson.M{"lineuserid": line_use_id}

	if err := t.db.Collection("member").FindOne(context.TODO(), options).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (t MemberRepo) UpdateRemainingBalance(line_use_id string, amount float64) error {

	filter := bson.M{"lineuserid": line_use_id}
	update := bson.M{
		"$set": bson.M{
			"remaining": amount,
		},
	}
	_, err := t.db.Collection("member").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
