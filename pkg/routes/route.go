package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/pkg/handlers"
	"github.com/tegehhat/helper/pkg/middleware"
)

func GetRoute() *gin.Engine {
	r := gin.Default()

	//Авторизация
	r.POST("/auth", handlers.Login)
	r.DELETE("/auth", handlers.Logout)

	//Проверка авторизации
	r.Use(middleware.RequireAuth())

	//Направления
	r.GET("/direction", handlers.GetDirection)
	r.GET("/direction/:id", handlers.GetDirection)
	r.PATCH("/direction/", handlers.UpdateDirection)
	r.DELETE("/direction/", handlers.DeleteDirection)

	//TODO: Все остальные справочники

	return r
}

func testf(c *gin.Context) {
	fmt.Print("test")
}
