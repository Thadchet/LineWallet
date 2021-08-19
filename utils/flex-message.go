package utils

import (
	"fmt"
	"line-wallet/constants"
	"line-wallet/models"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func TransactionCompleteFlex(amount string, category string, memo string, total float64, remaining float64) *linebot.FlexMessage {
	day := time.Now().Day()
	month := time.Now().Month()
	year := time.Now().Year()
	if memo == "" {
		memo = "-"
	}

	totalStr := strconv.FormatFloat(total, 'f', 2, 64)
	remainingStr := strconv.FormatFloat(remaining, 'f', 2, 64)
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
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Margin: linebot.FlexComponentMarginTypeMd,
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeHorizontal,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Text:   "Total Txn",
								Size:   linebot.FlexTextSizeTypeMd,
								Weight: linebot.FlexTextWeightTypeBold,
								Color:  constants.GreenColor,
							},
							&linebot.TextComponent{
								Text:   totalStr,
								Size:   linebot.FlexTextSizeTypeMd,
								Align:  linebot.FlexComponentAlignTypeEnd,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
					},
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeHorizontal,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Text:   "Remaining",
								Size:   linebot.FlexTextSizeTypeMd,
								Weight: linebot.FlexTextWeightTypeBold,
								Color:  constants.GreenColor,
							},
							&linebot.TextComponent{
								Text:   remainingStr,
								Size:   linebot.FlexTextSizeTypeMd,
								Align:  linebot.FlexComponentAlignTypeEnd,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
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

func SummaryCurrentMonth(txns []models.Transaction, incomes []models.Income, total float64, remaining float64) *linebot.FlexMessage {
	month := time.Now().Month()
	year := time.Now().Year()

	totalStr := strconv.FormatFloat(total, 'f', 2, 64)
	remainingStr := strconv.FormatFloat(remaining, 'f', 2, 64)

	var itemsTxn []linebot.FlexComponent
	for index, income := range incomes {
		item := &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   fmt.Sprintf("%v. %v", index+1, income.Type),
					Weight: linebot.FlexTextWeightTypeRegular,
					Color:  constants.GrayColor,
				},
				&linebot.TextComponent{
					Text:   income.Amount,
					Weight: linebot.FlexTextWeightTypeRegular,
					Align:  linebot.FlexComponentAlignTypeEnd,
					Color:  constants.GrayColor,
				},
			},
		}
		itemsTxn = append(itemsTxn, item)
	}

	for index, txn := range txns {
		item := &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   fmt.Sprintf("%v. %v", index+1, txn.Category),
					Weight: linebot.FlexTextWeightTypeRegular,
					Color:  constants.GrayColor,
				},
				&linebot.TextComponent{
					Text:   txn.Amount,
					Weight: linebot.FlexTextWeightTypeRegular,
					Align:  linebot.FlexComponentAlignTypeEnd,
					Color:  constants.GrayColor,
				},
			},
		}
		itemsTxn = append(itemsTxn, item)
	}

	itemsTxn = append(itemsTxn, &linebot.SeparatorComponent{
		Margin: linebot.FlexComponentMarginTypeMd,
		Type:   linebot.FlexComponentTypeBox,
	})
	itemsTxn = append(itemsTxn, &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Margin: linebot.FlexComponentMarginTypeMd,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Total Txn",
						Size:   linebot.FlexTextSizeTypeMd,
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  constants.GreenColor,
					},
					&linebot.TextComponent{
						Text:   totalStr,
						Size:   linebot.FlexTextSizeTypeMd,
						Align:  linebot.FlexComponentAlignTypeEnd,
						Weight: linebot.FlexTextWeightTypeBold,
					},
				},
			}}})
	itemsTxn = append(itemsTxn, &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Text:   "Remaining",
				Size:   linebot.FlexTextSizeTypeMd,
				Weight: linebot.FlexTextWeightTypeBold,
				Color:  constants.GreenColor,
			},
			&linebot.TextComponent{
				Text:   remainingStr,
				Size:   linebot.FlexTextSizeTypeMd,
				Align:  linebot.FlexComponentAlignTypeEnd,
				Weight: linebot.FlexTextWeightTypeBold,
			},
		},
	},
	)

	body := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: itemsTxn,
	}

	// Make Header
	header := linebot.BoxComponent{
		Type:          linebot.FlexComponentTypeBox,
		Layout:        linebot.FlexBoxLayoutTypeVertical,
		Spacing:       linebot.FlexComponentSpacingTypeMd,
		PaddingBottom: linebot.FlexComponentPaddingTypeNone,
		Contents: []linebot.FlexComponent{
			// &linebot.TextComponent{
			// 	Text:   "BK Wallet",
			// 	Weight: linebot.FlexTextWeightTypeBold,
			// 	Color:  constants.GreenColor,
			// },
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "Summary",
						Size:   linebot.FlexTextSizeTypeXxl,
						Weight: linebot.FlexTextWeightTypeBold,
					},
					&linebot.TextComponent{
						Size:   linebot.FlexTextSizeTypeSm,
						Weight: linebot.FlexTextWeightTypeRegular,
						Color:  constants.GrayColor,
						Text:   fmt.Sprintf("%v %v", month, year),
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
	flexMessage := linebot.NewFlexMessage("Summary", &bubble)
	return flexMessage
}

func IncomeCompleteFlex(amount, month, memo string, total float64) *linebot.FlexMessage {

	year := time.Now().Year()
	if memo == "" {
		memo = "-"
	}

	totalStr := strconv.FormatFloat(total, 'f', 2, 64)
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
						Text:   "Income",
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
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Margin: linebot.FlexComponentMarginTypeMd,
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeHorizontal,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Text:   "Total Txn",
								Size:   linebot.FlexTextSizeTypeMd,
								Weight: linebot.FlexTextWeightTypeBold,
								Color:  constants.GreenColor,
							},
							&linebot.TextComponent{
								Text:   totalStr,
								Size:   linebot.FlexTextSizeTypeMd,
								Align:  linebot.FlexComponentAlignTypeEnd,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
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
						Text:   "Income",
						Size:   linebot.FlexTextSizeTypeXxl,
						Weight: linebot.FlexTextWeightTypeBold,
					},
					&linebot.TextComponent{
						Size:   linebot.FlexTextSizeTypeSm,
						Weight: linebot.FlexTextWeightTypeRegular,
						Color:  constants.GrayColor,
						Text:   fmt.Sprintf("%v %v", month, year),
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
	flexMessage := linebot.NewFlexMessage("Income complete", &bubble)
	return flexMessage
}
