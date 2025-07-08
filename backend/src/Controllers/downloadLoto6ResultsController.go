package Controllers

import (
	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func DownloadLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto6Results()
		return c.File("Loto6Results.csv")
	}
}
