package repository

import "line-wallet/config"

type Repository struct {
	Transaction ITransactionRepo
}

func NewRepo(conf config.Config) Repository {
	var repo Repository
	repo.Transaction = TransactionRepo{conf.MongoDB.GetDatabase()}
	return repo
}
