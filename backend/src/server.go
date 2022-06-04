package main

import (
	controller "./Controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Static("/", "../../frontend/dist")

	//MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	e.GET("/getLoto6Results", controller.GetLoto6Results())

	e.Logger.Fatal(e.Start(":8000"))
}
