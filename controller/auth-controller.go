package controller

import (
	"fmt"
	"golang-br/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	res := &model.Response{
		Status:  "Success",
		Message: "passing code",
	}

	return c.JSON(http.StatusOK, res)
}

func Upstream(c echo.Context) error {
	name := c.Param("name")
	res := &model.Response{
		Status:  "Success",
		Message: fmt.Sprintf("hellow %v", name),
	}
	return c.JSON(http.StatusOK, res)
}
