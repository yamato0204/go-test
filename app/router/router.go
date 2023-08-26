package router

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/go-test.git/app/controller"
)

func Route(uc controller.IUserController) *echo.Echo{
	e := echo.New()

	e.POST("/create", uc.UserCreate)

	return e
}