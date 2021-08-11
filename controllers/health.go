package controller

import (
	"line-wallet/config"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	cv config.Config
}

func NewHealthHandler(conf config.Config) HealthHandler {
	return HealthHandler{
		cv: conf,
	}
}

func (h HealthHandler) Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Working !!",
	})
}
