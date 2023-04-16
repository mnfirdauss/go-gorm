package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/mnfirdauss/go-gorm/conf"
	"github.com/mnfirdauss/go-gorm/model"
	"github.com/mnfirdauss/go-gorm/service"
)

type Controller struct {
}

func (m *Controller) GetUser(c echo.Context) error {
	// user := c.Get("user").(model.User)
	// log.Printf("user data: %+v/n", user)

	var users []model.User

	conf.DBMysql.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (m *Controller) CreateUser(c echo.Context) error {
	data := map[string]interface{}{
		"message": "fail",
	}
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	err = service.GetUserRepository().CreateUser(&user)
	if err != nil {
		return err
	}

	data["message"] = "success"
	return c.JSON(http.StatusOK, data)
}

func CreateBatchUser(c echo.Context) error {

	data := echo.Map{
		"message": "success",
	}

	var users []model.User
	err := c.Bind(&users)
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	result := conf.DBMysql.Create(&users)
	if result.Error != nil {
		data["message"] = result.Error
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}

func DeleteUser(c echo.Context) error {
	var user model.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	conf.DBMysql.Unscoped().Delete(&user, id)

	data := &echo.Map{
		"message": "success",
	}
	return c.JSON(http.StatusOK, data)
}

func UpdateUser(c echo.Context) error {

	data := echo.Map{
		"message": "success",
	}

	var user model.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	username := c.Request().PostFormValue("username")

	result := conf.DBMysql.Model(&user).Where("id = ?", id).Update("username", username)
	if result.Error != nil {
		data["message"] = result.Error
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}
