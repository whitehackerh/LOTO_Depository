package Controllers

import (
	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func DownloadLoto6LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto6LatelyStatistics()
		return c.File("Loto6LatelyStatistics.csv")
	}
}
