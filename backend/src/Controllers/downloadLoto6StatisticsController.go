package Controllers

import (
	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func DownloadLoto6Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto6Statistics()
		return c.File("Loto6Statistics.csv")
	}
}
