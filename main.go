package main

import (
	"github.com/joho/godotenv"
	"github.com/mnfirdauss/go-gorm/conf"
	"github.com/mnfirdauss/go-gorm/model"
	"github.com/mnfirdauss/go-gorm/route"
	"github.com/mnfirdauss/go-gorm/utils"
)

func init() {
	godotenv.Load(".env")
	conf.InitDB()
	model.InitMigrate(asd)
}

func main() {
	route := route.StartRoute()

	utils.CreateToken(123, "daus")

	route.Start(":8080")
}
