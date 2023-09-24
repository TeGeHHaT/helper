package app

import (
	"log"

	"github.com/tegehhat/helper/pkg/database"
	"github.com/tegehhat/helper/pkg/routes"
)

func Run() {
	// Запускаем БД
	err := database.OpenConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseConnect()

	// Загружаем приложение
	routes.GetRoute().Run(":8083")

}
