package utils

import (
	"fmt"
	"line-wallet/config"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ILineService interface {
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	GetProfile(userID string) *linebot.GetProfileCall
	ReplyMessage(replyToken string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error)
	PushMessage(to string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error)
}

type LineService struct {
	Bot *linebot.Client
}

func NewLineService(conf config.Config) LineService {
	var lineService LineService
	bot, err := linebot.New(conf.GetChannelSecret(), conf.GetChannelAccessToken())
	if err != nil {
		fmt.Println(err)
	}
	lineService.Bot = bot
	fmt.Println("Initiate Line Service Successful")
	botInfo, _ := bot.GetBotInfo().Do()
	fmt.Println("Bot Name ==> ", botInfo.DisplayName)
	return lineService
}

func (s LineService) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return s.Bot.ParseRequest(r)
}

func (s LineService) GetProfile(userID string) *linebot.GetProfileCall {
	return s.Bot.GetProfile(userID)
}

func (s LineService) ReplyMessage(replyToken string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error) {
	res, err := s.Bot.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s LineService) PushMessage(to string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error) {
	res, err := s.Bot.PushMessage(to, messages...).Do()
	if err != nil {
		return nil, err
	}
	return res, nil
}
