package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/models"
)

func Login(c *gin.Context) {
	var user models.UserAuth
	json.NewDecoder(c.Request.Body).Decode(&user)

	key := time.Now().Format("2006-01-02 15:04:05")
	keyHash := md5.Sum([]byte(key))
	k := hex.EncodeToString(keyHash[:])

	c.Request.Header.Add("X-Token", k)

	c.JSON(http.StatusInternalServerError, k)
}

func checkLogin(u models.UserAuth) {

}
