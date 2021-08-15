package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//middleware
		fmt.Println("Middleware")
	}
}

type Header struct {
	LineUserId string `header:"line_user_id"`
}

func AuthMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var h Header
		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithStatusJSON(500, err)
		}
		if h.LineUserId == "" {
			fmt.Println("Line user id not found")
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorize missing line_user_id",
			})
			return
		}

	}
}

func HandleAuthLevel(authLevel int, endpoint gin.HandlerFunc) []gin.HandlerFunc {
	var endPoints []gin.HandlerFunc
	switch authLevel {
	case 0:
		break
	case 1:
		endPoints = append(endPoints, AuthMember())
	}
	endPoints = append(endPoints, endpoint)
	return endPoints
}
