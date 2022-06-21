package Controller

import (
	service "../Service"
	"github.com/labstack/echo"
)

func DownloadLoto7Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto7Results()
		return c.File("Loto7Results.csv")
	}
}
