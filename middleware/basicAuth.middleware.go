package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mnfirdauss/go-gorm/conf"
	"github.com/mnfirdauss/go-gorm/model"
	"github.com/mnfirdauss/go-gorm/utils"
)

func BasicAuthMiddleware(e *echo.Echo) {
	e.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		// id := ctx.Param("id")
		var user model.User
		conf.DBMysql.First(&user, "username = ?", username)
		// conf.DBMysql.First(&user).Where("username = ?", username)
		hashPassword, err := utils.HashPassword(password)
		if err != nil {
			return false, err
		}

		if username != user.Username {
			err = utils.ComparePassword(hashPassword, user.Password)
			if err != nil {
				log.Println("error compare password", hashPassword, user.Password)
				return false, err
			}
		}

		ctx.Set("user", user)
		return true, nil
	}))
}

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
