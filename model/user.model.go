package model

import (
	"github.com/mnfirdauss/go-gorm/conf"
	"github.com/mnfirdauss/go-gorm/utils"
	"gorm.io/gorm"
)

func init() {
	conf.DBMysql.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword

	return
}
