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

func Logout(c *gin.Context) {
	token := c.GetHeader("X-Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, "Missing Token")
		return
	}

	database.DB.Exec("update t_session_user set expires_at = now() WHERE token = $1", token)

	c.Request.Header.Del("X-Token")

	c.JSON(http.StatusOK, "")
}

// Вход /login
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

// Добавляем пользователя
func addUser(u *models.UserAuth) error {
	var result int
	err := database.DB.QueryRow("select 1 from t_user where user_login = $1", u.Name).Scan(&result)
	if err != nil {
		return err
	}

	hashedInputPassword, err := passwordHash(u.Password)
	if err != nil {
		return err
	}

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

// Ищем пользователя
func findUser(u *models.UserAuth, c *gin.Context) (int, error) {
	var result int
	err := database.DB.QueryRow("select 1 from t_user where user_login = $1", u.Name).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = database.DB.QueryRow("select 1 from t_session_user where user_name = $1", u.Name).Scan(&result)
	if err != nil {
		// Если пользователь не найден - добавляем в таблицу сессий
		if err == sql.ErrNoRows {
			err := addUser(u)
			if err != nil {
				return 0, err
			}
		}
		return 0, err
	}

	hashedInputPassword, err := passwordHash(u.Password)
	if err != nil {
		return 0, err
	}

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

// Создаём сессию
func createSession(userId int, c *gin.Context) (string, error) {

	token := tokenHash()

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

// Получаем сессию
func GetSession(token string) (*models.Session, error) {
	var session models.Session

	err := database.DB.QueryRow("SELECT token, user_id, created_at, expires_at FROM t_session WHERE token = $1", token).Scan(
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

func passwordHash(password string) (string, error) {
	if password == "" {
		return "", errors.New("Incorrect password")
	}

	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func tokenHash() string {
	key := time.Now().Format("2006-01-02 15:04:05")
	keyHash := md5.Sum([]byte(key))
	return hex.EncodeToString(keyHash[:])
}
