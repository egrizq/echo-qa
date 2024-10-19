package main

import (
	"golang-br/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/", controller.Login)
	app.GET("/hello/:name", controller.Upstream)
	app.Logger.Fatal(app.Start(":3000"))
}
