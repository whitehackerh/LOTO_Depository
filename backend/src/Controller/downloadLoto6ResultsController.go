package Controller

import (
	service "../Service"
	"github.com/labstack/echo"
)

func DownloadLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto6Results()
		return c.File("Loto6Results.csv")
	}
}
