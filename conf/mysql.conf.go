package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBMysql *gorm.DB

func init() {
	initDB()
}

func initDB() {
	var err error
	DBMysql, err = gorm.Open(mysql.Open("root:F_irdaus31@tcp(docker.for.mac.localhost:3306)/alterra?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
