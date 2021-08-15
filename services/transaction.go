package services

import (
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/repository"
	"time"
)

type TransactionService struct {
	Conf config.Config
	Repo repository.Repository
}

type ITransactionService interface {
	Ping() string
	AddTransaction(req models.AddTransactionRequest, member models.Member) error
	GetTreansactions(line_user_id string) ([]models.Transaction, error)
}

func (t TransactionService) Ping() string {
	if err := t.Repo.Transaction.Insert(); err != nil {
		return err.Error()
	}
	return "Pong"
}

func (t TransactionService) AddTransaction(req models.AddTransactionRequest, member models.Member) error {
	transaction := models.Transaction{
		Amount:     req.Amount,
		Category:   req.Category,
		Memo:       req.Memo,
		LineUserId: member.LineUserID,
		CreatedAt:  time.Now(),
	}
	if err := t.Repo.Transaction.InsertTransaction(transaction); err != nil {
		return err
	}
	return nil
}

func (t TransactionService) GetTreansactions(line_user_id string) ([]models.Transaction, error) {
	transactions, err := t.Repo.Transaction.GetTransactions(line_user_id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
