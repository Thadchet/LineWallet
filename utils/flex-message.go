package utils

import (
	"fmt"
	"line-wallet/constants"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func TransactionCompleteFlex(amount string, category string, memo string) *linebot.FlexMessage {
	day := time.Now().Day()
	month := time.Now().Month()
	year := time.Now().Year()

	body := linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Amount",
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  constants.GrayColor,
					},
					&linebot.TextComponent{
						Text:   fmt.Sprintf("%v %v", amount, "Baht"),
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeEnd,
						Color:  constants.GrayColor,
					},
				},
			},
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Category",
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  constants.GrayColor,
					},
					&linebot.TextComponent{
						Text:   category,
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeEnd,
						Color:  constants.GrayColor,
					},
				},
			},
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Memo",
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  constants.GrayColor,
					},
					&linebot.TextComponent{
						Text:   memo,
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeEnd,
						Color:  constants.GrayColor,
					},
				},
			},
			&linebot.SeparatorComponent{
				Margin: linebot.FlexComponentMarginTypeMd,
				Type:   linebot.FlexComponentTypeBox,
			},
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Margin: linebot.FlexComponentMarginTypeMd,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Total",
						Size:   linebot.FlexTextSizeTypeMd,
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  constants.GreenColor,
					},
					&linebot.TextComponent{
						Text:   "10000",
						Size:   linebot.FlexTextSizeTypeMd,
						Align:  linebot.FlexComponentAlignTypeEnd,
						Weight: linebot.FlexTextWeightTypeBold,
					},
				},
			},
		},
	}

	// Make Header
	header := linebot.BoxComponent{
		Type:          linebot.FlexComponentTypeBox,
		Layout:        linebot.FlexBoxLayoutTypeVertical,
		Spacing:       linebot.FlexComponentSpacingTypeMd,
		PaddingBottom: linebot.FlexComponentPaddingTypeNone,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Text:   "BK Wallet",
				Weight: linebot.FlexTextWeightTypeBold,
				Color:  constants.GreenColor,
			},
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Transaction",
						Size:   linebot.FlexTextSizeTypeXxl,
						Weight: linebot.FlexTextWeightTypeBold,
					},
					&linebot.TextComponent{
						Size:   linebot.FlexTextSizeTypeSm,
						Weight: linebot.FlexTextWeightTypeRegular,
						Color:  constants.GrayColor,
						Text:   fmt.Sprintf("%v %v %v", day, month, year),
					},
				},
			},
			&linebot.SeparatorComponent{
				Type: linebot.FlexComponentTypeBox,
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
	flexMessage := linebot.NewFlexMessage("Transaction complete", &bubble)
	return flexMessage
}
