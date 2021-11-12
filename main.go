package main

import (
	"fmt"
	"playground/play-dcoker-be4/controllers"
	"playground/play-dcoker-be4/models"

	"github.com/labstack/echo/v4"
)

func main() {
	// fmt.Println("Hello World")

	var userModel models.UserModel

	userModel = models.NewUserModel()

	e := echo.New()

	e.GET("/", controllers.Hello)

	userController := controllers.NewUserController(userModel)

	e.GET("/users", userController.GetAll)
	e.POST("/users", userController.Insert)

	if err := e.Start(fmt.Sprintf(":%d", 8080)); err != nil {
		panic(err)
	}

}
