package controller

import (
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/services"
	"line-wallet/utils"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
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
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := t.transactionService.AddTransaction(req, *member); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	leftBtn := linebot.NewMessageAction("Yes", "Yes clicked")
	rightBtn := linebot.NewMessageAction("No", "No clicked")
	template := linebot.NewConfirmTemplate("Are you John wick?", leftBtn, rightBtn)

	message := linebot.NewTemplateMessage("Confirm Box.", template)
	t.linebotService.PushMessage(member.LineUserID, message)

	c.JSON(200, gin.H{
		"message": req,
	})
}
