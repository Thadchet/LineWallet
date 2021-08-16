package utils

import (
	"fmt"
)

type PlayGroundService struct {
	LinebotService ILineService
}

func (s PlayGroundService) HandlePlayground(command string, replyToken string) {
	switch command {
	case "#a":
		flexMessage := TransactionCompleteFlex("3000", "Game", "Hello",400,300000)

		// Reply Message
		_, err := s.LinebotService.ReplyMessage(replyToken, flexMessage)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
