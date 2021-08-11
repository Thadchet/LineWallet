package routers

import (
	"fmt"
	"line-wallet/config"
	controller "line-wallet/controllers"
	"line-wallet/repository"
	"line-wallet/services"

	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
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
	v1 := router.Group("v1")
	
	repo := repository.NewRepo(conf)
	handles := InitHandlers(conf, repo)
	routes := newRoute(handles)

	for _, ro := range routes {
		v1.Handle(ro.Method, ro.Pattern, ro.Endpoint)
	}

	return router
}

type Handlers struct {
	Health      controller.HealthHandler
	Transaction controller.TransactionHandler
}

func InitHandlers(conf config.Config, repo repository.Repository) Handlers {
	var handlers Handlers
	service := services.NewService(conf, repo)
	handlers.Health = controller.NewHealthHandler(conf)
	handlers.Transaction = controller.NewTransactionHandler(conf, service)

	return handlers
}

func newRoute(handler Handlers) []route {
	return []route{
		{
			Name:     "health check",
			Method:   "GET",
			Pattern:  "/healthcheck",
			Endpoint: handler.Health.Healthcheck,
		},
		{
			Name:     "ping service",
			Method:   "GET",
			Pattern:  "/ping",
			Endpoint: handler.Transaction.PingTransactionService,
		},
	}
}
