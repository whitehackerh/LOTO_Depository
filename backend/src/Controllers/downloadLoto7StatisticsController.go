package Controllers

import (
	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func DownloadLoto7Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto7Statistics()
		return c.File("Loto7Statistics.csv")
	}
}
