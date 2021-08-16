package services

import (
	"fmt"
	"line-wallet/config"
	"line-wallet/constants"
	"line-wallet/models"
	"line-wallet/repository"
	"line-wallet/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type WebhookService struct {
	Conf              config.Config
	Repo              repository.Repository
	LinebotService    utils.ILineService
	PlayGroundService utils.PlayGroundService
}

type IWebhookService interface {
	ReplyMessageAction(linebot *linebot.Client, replyToken string)
	Follow(replyToken string, userId string) error
	HandleTextMessage(replyToken string, message *linebot.TextMessage)
}

func (w WebhookService) HandleTextMessage(replyToken string, message *linebot.TextMessage) {
	command := message.Text
	switch command {
	case "Bk", "bk", "BK":
		var ReplyButtons []*linebot.QuickReplyButton
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewURIAction("💸 เพิ่มรายการ", constants.HomePage)))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewMessageAction("💰 เพิ่มรายรับ", "Hello")))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewURIAction("⚙️ แก้ไขรายการ", constants.EditPage)))
		ReplyButtons = append(ReplyButtons, linebot.NewQuickReplyButton("", linebot.NewMessageAction("📊 สรุปรายเดือน", "Hello")))

		replyMessage := linebot.NewTextMessage(constants.ReplyMessage).WithQuickReplies(linebot.NewQuickReplyItems(ReplyButtons...))
		_, err := w.LinebotService.ReplyMessage(replyToken, replyMessage)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "#a":
		fmt.Println("Case Playground")
		w.PlayGroundService.HandlePlayground("#a", replyToken)
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
		fmt.Println(err.Error())
		return err
	}
	return nil
}
