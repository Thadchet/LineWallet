package services

import (
	"fmt"
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/repository"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionService struct {
	Conf config.Config
	Repo repository.Repository
}

type ITransactionService interface {
	Ping() string
	AddTransaction(req models.AddTransactionRequest, member *models.Member) error
	GetTreansactions(line_user_id string) ([]models.Transaction, error)
	GetTransactionByID(ID string) (*models.Transaction, error)
	EditTransactionByID(req models.AddTransactionRequest, id string, member models.Member) (*models.Transaction, error)
	AddIncome(req models.Income, member models.Member) error
}

func (t TransactionService) Ping() string {
	if err := t.Repo.Transaction.Insert(); err != nil {
		return err.Error()
	}
	return "Pong"
}

func (t TransactionService) AddTransaction(req models.AddTransactionRequest, member *models.Member) error {
	txnID := primitive.NewObjectID()
	transaction := models.Transaction{
		ID:         txnID,
		Amount:     req.Amount,
		Category:   req.Category,
		Memo:       req.Memo,
		LineUserId: member.LineUserID,
		CreatedAt:  time.Now(),
	}
	if err := t.Repo.Transaction.InsertTransaction(transaction); err != nil {
		return err
	}
	amount, _ := strconv.Atoi(req.Amount)
	remaining := member.GetRemaining() - amount
	if err := t.Repo.Member.UpdateRemainingBalance(member.LineUserID, remaining); err != nil {
		return err
	}
	member.UpdateRemaining(req.Amount)

	return nil
}

func (t TransactionService) GetTreansactions(line_user_id string) ([]models.Transaction, error) {
	transactions, err := t.Repo.Transaction.GetTransactions(line_user_id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t TransactionService) GetTransactionByID(ID string) (*models.Transaction, error) {
	transaction, err := t.Repo.Transaction.GetTransactionByID(ID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t TransactionService) EditTransactionByID(req models.AddTransactionRequest, id string, member models.Member) (*models.Transaction, error) {

	result, err := t.Repo.Transaction.UpdateTransactionByID(req.Amount, req.Category, req.Memo, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return nil, nil
}

func (t TransactionService) AddIncome(req models.Income, member models.Member) error {
	txnID := primitive.NewObjectID()
	transaction := models.Income{
		ID:         txnID,
		Amount:     req.Amount,
		Month:      req.Month,
		Memo:       req.Memo,
		LineUserId: member.LineUserID,
		CreatedAt:  time.Now(),
	}
	if err := t.Repo.Transaction.InsertTransaction(transaction); err != nil {
		return err
	}
	amount, _ := strconv.Atoi(req.Amount)
	if err := t.Repo.Member.UpdateRemainingBalance(member.LineUserID, amount); err != nil {
		return err
	}

	return nil
}
