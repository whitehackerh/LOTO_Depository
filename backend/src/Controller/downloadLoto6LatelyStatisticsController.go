package Controller

import (
	service "../Service"
	"github.com/labstack/echo"
)

func DownloadLoto6LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto6LatelyStatistics()
		return c.File("Loto6LatelyStatistics.csv")
	}
}
