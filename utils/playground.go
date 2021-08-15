package utils

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type PlayGroundService struct {
	LinebotService ILineService
}

func (s PlayGroundService) HandlePlayground(command string, replyToken string) {
	switch command {
	case "#a":
		fmt.Println("Case1")
		// Make Contents
		var contents []linebot.FlexComponent
		text := linebot.TextComponent{
			Type:   linebot.FlexComponentTypeText,
			Text:   "Brown Cafe",
			Weight: "bold",
			Size:   linebot.FlexTextSizeTypeXl,
		}
		contents = append(contents, &text)

		// Make Body
		body := linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: contents,
		}
		// Make Header
		header := linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			// Spacing: linebot.FlexComponentSpacingTypeMd,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:  "BK Wallet",
					Color: "#56ee74",
				},
				&linebot.TextComponent{
					Text:   "Transaction",
					Size:   linebot.FlexTextSizeTypeLg,
					Weight: linebot.FlexTextWeightTypeBold,
				},
				&linebot.TextComponent{
					Text: "Name",
				},
			},
		}
		// Build Container
		bubble := linebot.BubbleContainer{
			Type:   linebot.FlexContainerTypeBubble,
			Header: &header,
			Body:   &body,
		}
		// New Flex Message
		flexMessage := linebot.NewFlexMessage("FlexWithCode", &bubble)
		// Reply Message
		s.LinebotService.ReplyMessage(replyToken, flexMessage)
	}
}
