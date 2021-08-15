package repository

import "line-wallet/config"

type Repository struct {
	Transaction ITransactionRepo
	Member      IMemberRepo
}

func NewRepo(conf config.Config) Repository {
	var repo Repository
	repo.Transaction = TransactionRepo{conf.MongoDB.GetDatabase()}
	repo.Member = MemberRepo{conf.MongoDB.GetDatabase()}
	return repo
}
