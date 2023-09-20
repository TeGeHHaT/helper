package app

import (
	"log"

	"github.com/tegehhat/helper/pkg/database"
	"github.com/tegehhat/helper/pkg/handlers"
)

func Run() {
	// Запускаем БД
	err := database.OpenConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseConnect()

	// Загружаем приложение
	handlers.GetRoute().Run(":8083")

}
