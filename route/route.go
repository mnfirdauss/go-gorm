package route

import (
	"github.com/labstack/echo/v4"

	"github.com/mnfirdauss/go-gorm/controller"
	"github.com/mnfirdauss/go-gorm/middleware"
)

func StartRoute() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Halo Dunia! Mantap!")
	})
	middleware.LogMiddleware(e)

	userController := controller.Controller{}
	userGroup := e.Group("/users")
	userGroup.POST("/", userController.CreateUser)

	// middleware.BasicAuthMiddleware(e)
	userGroup.GET("/", userController.GetUser)

	userGroup.POST("/batch", controller.CreateBatchUser)

	userGroup.PUT("/:id", controller.UpdateUser)
	userGroup.DELETE("/:id", controller.DeleteUser)

	profileGroup := e.Group("/profile")
	// profileGroup.Use(echojwt.JWT([]byte("secret")))
	profileGroup.GET("/", controller.GetProfile)

	return e
}
