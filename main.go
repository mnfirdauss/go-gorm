package main

import (
	"github.com/mnfirdauss/go-gorm/route"
	"github.com/mnfirdauss/go-gorm/utils"
)

func main() {
	route := route.StartRoute()

	utils.CreateToken(123, "daus")

	route.Start(":8080")
}
