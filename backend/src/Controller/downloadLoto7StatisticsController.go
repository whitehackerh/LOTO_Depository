package Controller

import (
	service "../Service"
	"github.com/labstack/echo"
)

func DownloadLoto7Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto7Statistics()
		return c.File("Loto7Statistics.csv")
	}
}
