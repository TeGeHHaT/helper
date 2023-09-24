package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/pkg/handlers"
)

func GetRoute() *gin.Engine {
	r := gin.Default()

	r.POST("/login", handlers.Login)

	r.GET("/direction", handlers.GetDirection)
	r.GET("/direction/:id", handlers.GetDirection)
	r.PATCH("/direction/:id", handlers.UpdateDirection)

	return r
}

func testf(c *gin.Context) {
	fmt.Print("test")
}