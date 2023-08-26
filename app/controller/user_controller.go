package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/go-test.git/app/entities"
	"github.com/yamato0204/go-test.git/app/usecase"
)


type IUserController interface {
	UserCreate(c echo.Context) error
}


type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &UserController{uu}
}



func (cc UserController)UserCreate(c echo.Context) error {
	user := entities.User{}

	if err := c.Bind(&user); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	user.ID = uuid.New().String()
	userRes, err := cc.uu.Create(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusCreated, userRes)

}