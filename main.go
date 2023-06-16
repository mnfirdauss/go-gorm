package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mnfirdauss/go-gorm/conf"
	"github.com/mnfirdauss/go-gorm/model"
	"github.com/mnfirdauss/go-gorm/route"
	"github.com/mnfirdauss/go-gorm/utils"
)

func init() {
	godotenv.Load(".env")
	conf.InitDB()
	model.InitMigrate()
}

func main() {
	route := route.StartRoute()
	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	utils.CreateToken(123, "daus")

	if err := route.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
