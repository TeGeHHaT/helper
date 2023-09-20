package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/login", testf)

	r.GET("/direction", GetDirection)
	r.GET("/direction/:id", GetDirection)
	r.PATCH("/direction", UpdateDirection)

	return r
}

func testf(c *gin.Context) {
	fmt.Print("test")
}
