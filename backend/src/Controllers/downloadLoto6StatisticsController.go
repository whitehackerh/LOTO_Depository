package Controllers

import (
	service "../Services"
	"github.com/labstack/echo"
)

func DownloadLoto6Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		service.DownloadLoto6Statistics()
		return c.File("Loto6Statistics.csv")
	}
}
