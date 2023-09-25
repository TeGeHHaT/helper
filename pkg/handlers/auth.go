package handlers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/models"
	"github.com/tegehhat/helper/pkg/database"
)

func Registation(c *gin.Context) {
	var user models.UserAuth
	json.NewDecoder(c.Request.Body).Decode(&user)

	err := addUser(&user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var user models.UserAuth
	json.NewDecoder(c.Request.Body).Decode(&user)

	userId, err := findUser(&user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	token, err := createSession(userId, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

func addUser(u *models.UserAuth, c *gin.Context) error {
	var result int
	err := database.DB.QueryRow("select 1 from t_user where user_login = $1", u.Name).Scan(&result)
	if err != nil {
		return err
	}

	hasher := md5.New()
	hasher.Write([]byte(u.Password))
	hashedInputPassword := hex.EncodeToString(hasher.Sum(nil))

	_, err = database.DB.Exec(
		"INSERT INTO t_session_user (user_name, password) VALUES ($1, $2)",
		u.Name,
		hashedInputPassword,
	)
	if err != nil {
		return err
	}

	return nil

}

func findUser(u *models.UserAuth, c *gin.Context) (int, error) {
	var result int
	err := database.DB.QueryRow("select 1 from t_user where user_login = $1", u.Name).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = database.DB.QueryRow("select 1 from t_session_user where user_name = $1", u.Name).Scan(&result)
	if err != nil {
		return 0, err
	}

	hasher := md5.New()
	hasher.Write([]byte(u.Password))
	hashedInputPassword := hex.EncodeToString(hasher.Sum(nil))

	var userHashedPassword string
	var userId int

	err = database.DB.QueryRow("SELECT id, password FROM t_session_user WHERE user_name = $1", u.Name).Scan(&userId, &userHashedPassword)
	if err != nil {
		return 0, err
	}

	if hashedInputPassword != userHashedPassword {
		return 0, errors.New("Incorrect password")
	}

	return userId, nil

}

func createSession(userId int, c *gin.Context) (string, error) {

	key := time.Now().Format("2006-01-02 15:04:05")
	keyHash := md5.Sum([]byte(key))
	token := hex.EncodeToString(keyHash[:])

	c.Request.Header.Add("X-Token", token)

	createdAt := time.Now()
	expiresAt := createdAt.Add(time.Hour * time.Duration(24))

	_, err := database.DB.Exec(
		"INSERT INTO t_session (token, user_id, created_at, expires_at) VALUES ($1, $2, $3, $4)",
		token,
		userId,
		createdAt,
		expiresAt,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetSession(token string) (*models.Session, error) {
	var session models.Session

	err := database.DB.QueryRow("SELECT token, user_name, created_at, expires_at FROM t_session WHERE token = $1", token).Scan(
		&session.Token,
		&session.User,
		&session.CreatedAt,
		&session.ExpiresAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Сессия не найдена
		}
		return nil, err
	}

	return &session, nil
}
