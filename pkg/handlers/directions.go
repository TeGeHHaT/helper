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
	var directionInsUpdParams models.DirectionInsUpdParams

	json.NewDecoder(c.Request.Body).Decode(&directionInsUpdParams)

	jsonParams, err := json.Marshal(directionInsUpdParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	var buf []byte
	err = database.DB.QueryRow(`SELECT * FROM public.fn_direction_ins_upd($1::jsonb)`, string(jsonParams)).Scan(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	res := models.Direction{}

	err = json.Unmarshal(buf, &res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteDirection(c *gin.Context) {
	var directionDelParams models.DirectionDelParams

	json.NewDecoder(c.Request.Body).Decode(&directionDelParams)

	jsonParams, err := json.Marshal(directionDelParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	database.DB.Exec(`SELECT * FROM public.fn_direction_del($1::jsonb)`, string(jsonParams))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}
}
