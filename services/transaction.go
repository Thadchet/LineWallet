package services

import (
	"line-wallet/config"
	"line-wallet/repository"
)

type TransactionService struct {
	Conf config.Config
	Repo repository.Repository
}

type ITransactionService interface {
	Ping() string
}

func NewTransactionService(conf config.Config) TransactionService {
	return TransactionService{
		Conf: conf,
	}
}

func (t TransactionService) Ping() string {
	if err := t.Repo.Transaction.Insert(); err != nil {
		return err.Error()
	}
	return "Pong"
}
