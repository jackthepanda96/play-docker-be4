package controllers

import (
	"net/http"
	"playground/play-dcoker-be4/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	model models.UserModel
}

func NewUserController(m models.UserModel) UserController {
	return UserController{model: m}
}

func (uc *UserController) GetAll(c echo.Context) error {
	allUser, err := uc.model.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Cannot get users")
	}
	return c.JSON(http.StatusOK, allUser)
}

func (uc *UserController) Insert(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	insertUser, err := uc.model.Insert(user)
	if err != nil {
		log.Info(err)
		return c.JSON(http.StatusBadRequest, "invalid user data")
	}

	return c.JSON(http.StatusOK, insertUser)
}
