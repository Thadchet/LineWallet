package services

import (
	"line-wallet/config"
	"line-wallet/repository"
	"line-wallet/utils"
)

type Services struct {
	TransactionService ITransactionService
	WebhookService     IWebhookService
	LinebotService     utils.ILineService
	MemberService      IMemberService
	PlayGroundService  utils.PlayGroundService
}

func NewService(conf config.Config, linebotService utils.ILineService, repo repository.Repository) Services {
	var services Services

	transactionService := TransactionService{
		Conf: conf,
		Repo: repo,
	}

	playGroundService := utils.PlayGroundService{
		LinebotService: linebotService,
	}

	webhookService := WebhookService{
		Conf:              conf,
		Repo:              repo,
		LinebotService:    linebotService,
		PlayGroundService: playGroundService,
	}

	memberService := MemberService{
		Conf: conf,
		Repo: repo,
	}

	services.TransactionService = transactionService
	services.WebhookService = webhookService
	services.LinebotService = linebotService
	services.MemberService = memberService
	services.PlayGroundService = playGroundService
	return services
}
