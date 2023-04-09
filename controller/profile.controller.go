package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mnfirdauss/go-gorm/conf"
	"github.com/mnfirdauss/go-gorm/model"

	"github.com/labstack/echo/v4"
)

type MyCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GetProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*MyCustomClaims)
	id := claims.ID

	var userModel model.User
	conf.DBMysql.First(&userModel, id)

	return c.JSON(http.StatusOK, userModel)
}
