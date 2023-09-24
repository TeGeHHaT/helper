package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/models"
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
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
	}

	directionParams := models.DirectionGetParams{Id: idInt}

	jsonParams, err := json.Marshal(directionParams)
	if err != nil {
		log.Println(err.Error())
	}

	var buf []byte
	err = database.DB.QueryRow(`SELECT * FROM public.fn_direction_get($1::jsonb)`, string(jsonParams)).Scan(&buf)
	if err != nil {
		log.Println(err.Error())
	}

	res := []Direction{}

	err = json.Unmarshal(buf, &res)
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, res)
}

func UpdateDirection(c *gin.Context) {
	test := `test`
	c.JSON(http.StatusOK, test)
}
