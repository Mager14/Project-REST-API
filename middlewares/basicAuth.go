package middlewares

import (
	"Project-REST-API/configs"
	"Project-REST-API/entities"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func BusicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, errors.New("bukan admin")
}

func GenerateToken(user entities.User) (string, error) {
	datas := jwt.MapClaims{}
	datas["id"] = user.ID
	datas["exp"] = time.Now().Add(time.Hour * 1).Unix() //1jam
	datas["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datas)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserId(e echo.Context) float64 {

	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		datas := user.Claims.(jwt.MapClaims)
		userId := datas["id"].(float64)
		return userId
	}

	return 0
}
