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
	e.POST("/setLoto6Results", controller.SetLoto6Results())
	e.POST("/determineLoto6Expectation", controller.DetermineLoto6Expectation())
	e.GET("/getLoto7Results", controller.GetLoto7Results())
	e.POST("/setLoto7Results", controller.SetLoto7Results())
	e.POST("/determineLoto7Expectation", controller.DetermineLoto7Expectation())

	e.Logger.Fatal(e.Start(":8000"))
}
