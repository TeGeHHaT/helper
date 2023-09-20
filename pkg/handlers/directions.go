package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/pkg/database"
)

type Direction struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	IsDisabled bool   `json:"is_disabled"`
}

func GetDirection(c *gin.Context) {
	var id int
	if c.Params.ByName("id") != "" {
		n, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		id = n
	}

	query := fmt.Sprintf("SELECT * FROM public.fn_direction_get('{\"id\": %v}')", id)

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var directionJSON string
	for rows.Next() {
		if err := rows.Scan(&directionJSON); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			log.Println("Ошибка в rows")
			log.Println(err)
			return
		}
	}

	c.JSON(http.StatusOK, directionJSON)
}

func UpdateDirection(c *gin.Context) {
	var direction Direction

	if err := c.BindJSON(&direction); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	query := fmt.Sprintf("SELECT * FROM public.fn_direction_get('{\"id\": %v}')", id)

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}
