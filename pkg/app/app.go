package app

import (
	"log"
	"time"

	"github.com/tegehhat/helper/pkg/database"
	"github.com/tegehhat/helper/pkg/routes"
)

func Run() {
	//Выставляем часовой пояс для приложения
	_, err := time.LoadLocation("Europe/Moscow")

	// Запускаем БД
	err = database.OpenConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseConnect()

	// Загружаем приложение
	routes.GetRoute().Run(":8083")

}
