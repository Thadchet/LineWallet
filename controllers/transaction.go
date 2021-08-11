package controller

import (
	"line-wallet/config"
	"line-wallet/services"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	cv                 config.Config
	transactionService services.ITransactionService
}

func NewTransactionHandler(conf config.Config, services services.Services) TransactionHandler {
	return TransactionHandler{
		cv:                 conf,
		transactionService: services.TransactionService,
	}
}

func (t TransactionHandler) PingTransactionService(c *gin.Context) {
	res := t.transactionService.Ping()
	c.JSON(200, gin.H{
		"message": res,
	})
}
