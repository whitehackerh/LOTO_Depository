package Controllers

import (
	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func DownloadLoto7LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto7LatelyStatistics()
		return c.File("Loto7LatelyStatistics.csv")
	}
}
