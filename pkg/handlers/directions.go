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

func GetDirection(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	directionParams := models.DirectionGetParams{Id: idInt}

	jsonParams, err := json.Marshal(directionParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	var buf []byte
	err = database.DB.QueryRow(`SELECT * FROM public.fn_direction_get($1::jsonb)`, string(jsonParams)).Scan(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	res := []models.Direction{}

	err = json.Unmarshal(buf, &res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateDirection(c *gin.Context) {
	test := `test`
	c.JSON(http.StatusOK, test)
}
