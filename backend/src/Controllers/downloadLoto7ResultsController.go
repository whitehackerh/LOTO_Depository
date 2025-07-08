package Controllers

import (
	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func DownloadLoto7Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto7Results()
		return c.File("Loto7Results.csv")
	}
}
