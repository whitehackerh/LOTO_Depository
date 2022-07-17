package Controller

import (
	service "../Service"
	"github.com/labstack/echo"
)

func DownloadLoto7LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto7LatelyStatistics()
		return c.File("Loto7LatelyStatistics.csv")
	}
}
