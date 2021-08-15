package routers

import (
	"fmt"
	"line-wallet/config"
	controller "line-wallet/controllers"
	"line-wallet/repository"
	"line-wallet/services"
	"line-wallet/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	AuthLevel   int
}

func Init(conf config.Config) {

	r := NewRouter(conf)
	address := fmt.Sprintf(":%v", conf.GetPort())
	r.Run(address)

}

func NewRouter(conf config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("line_user_id")
	router.Use(cors.New(config))

	api := router.Group("api")

	repo := repository.NewRepo(conf)
	linebot := utils.NewLineService(conf)
	handles := InitHandlers(conf, linebot, repo)
	routes := newRoute(handles)

	for _, ro := range routes {
		api.Handle(ro.Method, ro.Pattern, HandleAuthLevel(ro.AuthLevel, ro.Endpoint)...)
	}

	return router
}

type Handlers struct {
	Health      controller.HealthHandler
	Transaction controller.TransactionHandler
	Webhook     controller.WebhookHandler
}

func InitHandlers(conf config.Config, linebotService utils.LineService, repo repository.Repository) Handlers {
	var handlers Handlers
	service := services.NewService(conf, linebotService, repo)
	handlers.Health = controller.NewHealthHandler(conf)
	handlers.Transaction = controller.NewTransactionHandler(conf, service)
	handlers.Webhook = controller.NewWebhookHandler(conf, service)

	return handlers
}

func newRoute(handler Handlers) []route {
	return []route{
		{
			Name:        "health check",
			Description: "helth check",
			Method:      "GET",
			Pattern:     "/healthcheck",
			Endpoint:    handler.Health.Healthcheck,
			AuthLevel:   0,
		},
		{
			Name:        "ping service",
			Description: "ping service",
			Method:      "GET",
			Pattern:     "/ping",
			Endpoint:    handler.Transaction.PingTransactionService,
			AuthLevel:   0,
		},
		{
			Name:        "webhook",
			Description: "Web hook",
			Method:      "POST",
			Pattern:     "/webhook",
			Endpoint:    handler.Webhook.HandleWebhook,
			AuthLevel:   0,
		},
		{
			Name:        "Add Txn",
			Description: "Add txn",
			Method:      "POST",
			Pattern:     "/add_txn",
			Endpoint:    handler.Transaction.AddTransaction,
			AuthLevel:   1,
		},
		{
			Name:        "Get My Transactions",
			Description: "Get My Transactions",
			Method:      "GET",
			Pattern:     "/transactions",
			Endpoint:    handler.Transaction.GetTransactions,
			AuthLevel:   1,
		},
	}
}
