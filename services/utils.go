package services

import (
	"line-wallet/config"
	"line-wallet/repository"
)

type Services struct {
	TransactionService ITransactionService
}

func NewService(conf config.Config, repo repository.Repository) Services {
	var services Services

	transactionService := TransactionService{
		Conf: conf,
		Repo: repo,
	}
	
	services.TransactionService = transactionService
	return services
}
