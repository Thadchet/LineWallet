package utils

import (
	"line-wallet/models"
	"log"
)

type PlayGroundService struct {
	LinebotService ILineService
}

func (s PlayGroundService) HandlePlayground(command string, replyToken string) {
	switch command {
	case "#a":
		flexMessage := TransactionCompleteFlex("3000", "Game", "Hello", 400.00, 300000)

		// Reply Message
		_, err := s.LinebotService.ReplyMessage(replyToken, flexMessage)
		if err != nil {
			log.Println(err.Error())
		}
	case "#สรุปเดือนนี้":
		txn := []models.Transaction{
			{
				Amount:   "10",
				Memo:     "TT",
				Category: "Game",
				Type:     "txn",
			},
			{
				Amount:   "20",
				Memo:     "TT",
				Category: "Game",
				Type:     "txn",
			},
			{
				Amount:   "30",
				Memo:     "TT",
				Category: "Game",
				Type:     "txn",
			},
		}
		income := []models.Income{
			{
				Amount: "10000",
				Month:  "Aug",
				Type:   "income",
			},
		}
		flexMessage := SummaryCurrentMonth(txn, income, 400.00, 300000)
		// Reply Message
		_, err := s.LinebotService.ReplyMessage(replyToken, flexMessage)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
