package controller

import (
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/services"
	"line-wallet/utils"
	"log"

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

	if req.Amount == "" || req.Category == "" {
		log.Println("Missing field")
		c.JSON(400, gin.H{
			"data": "Missing field",
		})
		return
	}

	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := t.transactionService.AddTransaction(req, member); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": req,
	})
}

func (t TransactionHandler) GetTransactions(c *gin.Context) {

	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	res, err := t.transactionService.GetTreansactions(member.LineUserID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": res,
	})
}

func (t TransactionHandler) GetTransactionByID(c *gin.Context) {

	id := c.Param("id")
	_, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	res, err := t.transactionService.GetTransactionByID(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": &res,
	})
}

func (t TransactionHandler) EditTransactionByID(c *gin.Context) {

	id := c.Param("id")
	var req models.AddTransactionRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	res, err := t.transactionService.EditTransactionByID(req, id, *member)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": &res,
	})
}

func (t TransactionHandler) AddIncome(c *gin.Context) {
	var req models.Income
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if req.Amount == "" || req.Month == "" {
		log.Println("Missing field")
		c.JSON(400, gin.H{
			"data": "Missing field",
		})
		return
	}
	
	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := t.transactionService.AddIncome(req, *member); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": req,
	})
}

func (t TransactionHandler) GetIncomeByID(c *gin.Context) {
	id := c.Param("id")
	_, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	res, err := t.transactionService.GetIncomeByID(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": &res,
	})
}

func (t TransactionHandler) EditIncomeByID(c *gin.Context) {

	id := c.Param("id")
	var req models.Income
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	member, err := t.memberService.FindMemberByLineUserID(c.Request.Header["Line_user_id"][0])
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	res, err := t.transactionService.EditIncomeByID(req, id, *member)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": &res,
	})
}
