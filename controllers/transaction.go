package controller

import (
	"fmt"
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/services"
	"line-wallet/utils"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	cv                 config.Config
	transactionService services.ITransactionService
	memberService      services.IMemberService
	linebotService     utils.ILineService
}

func NewTransactionHandler(conf config.Config, services services.Services) TransactionHandler {
	return TransactionHandler{
		cv:                 conf,
		transactionService: services.TransactionService,
		memberService:      services.MemberService,
		linebotService:     services.LinebotService,
	}
}

func (t TransactionHandler) PingTransactionService(c *gin.Context) {
	res := t.transactionService.Ping()
	c.JSON(200, gin.H{
		"message": res,
	})
}

func (t TransactionHandler) AddTransaction(c *gin.Context) {
	var req models.AddTransactionRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := t.transactionService.AddTransaction(req, *member); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	flexMessage := utils.TransactionCompleteFlex(req.Amount, req.Category, req.Memo)
	_, err2 := t.linebotService.PushMessage(member.LineUserID, flexMessage)
	if err2 != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": req,
	})
}

func (t TransactionHandler) GetTransactions(c *gin.Context) {

	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	res, err := t.transactionService.GetTreansactions(member.LineUserID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": res,
	})
}
