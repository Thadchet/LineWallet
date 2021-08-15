package controller

import (
	"fmt"
	"line-wallet/config"
	"line-wallet/services"
	"line-wallet/utils"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type WebhookHandler struct {
	cv             config.Config
	webhookService services.IWebhookService
	linebotService utils.ILineService
}

func NewWebhookHandler(conf config.Config, services services.Services) WebhookHandler {
	return WebhookHandler{
		cv:             conf,
		webhookService: services.WebhookService,
		linebotService: services.LinebotService,
	}
}

func (t WebhookHandler) HandleWebhook(c *gin.Context) {

	events, _ := t.linebotService.ParseRequest(c.Request)
	for _, event := range events {
		fmt.Println("Type ==> ", event.Type)

		message := event.Message
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := message.(type) {
			case *linebot.TextMessage:
				replyToken := event.ReplyToken
				t.webhookService.HandleTextMessage(replyToken, message)
			case *linebot.ImageMessage:
				fmt.Println(message)
			}

		case linebot.EventTypeFollow:

			// TODO add member
			replyToken := event.ReplyToken
			userId := event.Source.UserID
			t.webhookService.Follow(replyToken, userId)
		}

	}
}
