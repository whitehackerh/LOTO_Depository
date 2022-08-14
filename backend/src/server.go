package main

import (
	controller "./Controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echomw "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Static("/", "../../frontend/dist")

	//MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType},
		AllowMethods: []string{echo.POST},
	}))

	// Routing
	e.POST("/login", controller.Login())
	e.GET("/getLoto6Results", controller.GetLoto6Results())
	e.POST("/setLoto6Results", controller.SetLoto6Results())
	e.POST("/determineLoto6Expectation", controller.DetermineLoto6Expectation())
	e.GET("/getLoto6Statistics", controller.GetLoto6Statistics())
	e.GET("/downloadLoto6Results", controller.DownloadLoto6Results())
	e.GET("/downloadLoto6Statistics", controller.DownloadLoto6Statistics())
	e.GET("/getLoto6LatelyStatistics", controller.GetLoto6LatelyStatistics())
	e.GET("/downloadLoto6LatelyStatistics", controller.DownloadLoto6LatelyStatistics())
	e.POST("/getLoto6ResultDetail", controller.GetLoto6ResultDetail())
	e.POST("/editLoto6Result", controller.EditLoto6Result())
	e.GET("/getNewestLoto6Result", controller.GetNewestLoto6Result())
	e.POST("/setLoto6Expectations", controller.SetLoto6Expectations())
	e.GET("/getLoto7Results", controller.GetLoto7Results())
	e.POST("/setLoto7Results", controller.SetLoto7Results())
	e.POST("/determineLoto7Expectation", controller.DetermineLoto7Expectation())
	e.GET("/getLoto7Statistics", controller.GetLoto7Statistics())
	e.GET("/downloadLoto7Results", controller.DownloadLoto7Results())
	e.GET("/downloadLoto7Statistics", controller.DownloadLoto7Statistics())
	e.GET("/getLoto7LatelyStatistics", controller.GetLoto7LatelyStatistics())
	e.GET("/downloadLoto7LatelyStatistics", controller.DownloadLoto7LatelyStatistics())
	e.POST("/getLoto7ResultDetail", controller.GetLoto7ResultDetail())
	e.POST("/editLoto7Result", controller.EditLoto7Result())
	e.GET("/getNewestLoto7Result", controller.GetNewestLoto7Result())
	e.POST("/createUser", controller.CreateUser())

	e.Logger.Fatal(e.Start(":8000"))
}
