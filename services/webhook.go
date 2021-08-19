package services

import (
	"fmt"
	"line-wallet/config"
	"line-wallet/constants"
	"line-wallet/models"
	"line-wallet/repository"
	"line-wallet/utils"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type WebhookService struct {
	Conf               config.Config
	Repo               repository.Repository
	LinebotService     utils.ILineService
	PlayGroundService  utils.PlayGroundService
	TransactionService ITransactionService
}

type IWebhookService interface {
	ReplyMessageAction(linebot *linebot.Client, replyToken string)
	Follow(replyToken string, userId string) error
	HandleTextMessage(replyToken string, message *linebot.TextMessage, event *linebot.Event)
}

func (w WebhookService) HandleTextMessage(replyToken string, message *linebot.TextMessage, event *linebot.Event) {
	command := message.Text
	switch command {
	case "Bk", "bk", "BK":
		var ReplyButtons []*linebot.QuickReplyButton
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewURIAction("💸 เพิ่มรายการ", constants.HomePage)))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewURIAction("💰 เพิ่มรายรับ", constants.AddIncomePage)))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewURIAction("⚙️ แก้ไขรายการ", constants.EditPage)))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewMessageAction("📊 สรุปเดือนนี้", "#สรุปเดือนนี้")))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewMessageAction("📊 สรุปรายเดือนทั้งหมด", "#สรุปรายเดือนทั้งหมด")))

		replyMessage := linebot.NewTextMessage(constants.ReplyMessage).WithQuickReplies(linebot.NewQuickReplyItems(ReplyButtons...))
		_, err := w.LinebotService.ReplyMessage(replyToken, replyMessage)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "#a":
		w.PlayGroundService.HandlePlayground("#a", replyToken)
	case "#สรุปเดือนนี้":
		w.TransactionService.SummaryCurrentMonth(replyToken, event.Source.UserID)
	default:
		break
	}

}

func (w WebhookService) ReplyMessageAction(bot *linebot.Client, replyToken string) {

	leftBtn := linebot.NewMessageAction("Yes", "Yes clicked")
	rightBtn := linebot.NewMessageAction("No", "No clicked")

	template := linebot.NewConfirmTemplate("Are you John wick?", leftBtn, rightBtn)

	message := linebot.NewTemplateMessage("Confirm Box.", template)
	bot.ReplyMessage(replyToken, message).Do()
}

func (w WebhookService) Follow(replyToken string, userId string) error {

	profileFromLine, err := w.LinebotService.GetProfile(userId).Do()
	if err != nil {
		return err
	}

	member := models.Member{
		Name:         profileFromLine.DisplayName,
		ProfileImage: profileFromLine.PictureURL,
		LineUserID:   profileFromLine.UserID,
	}
	if err := w.Repo.Member.CreateMember(member); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
