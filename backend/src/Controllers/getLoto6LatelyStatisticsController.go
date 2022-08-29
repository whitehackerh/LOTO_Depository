package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

func GetLoto6LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto6LatelyStatistics()
		return c.JSON(http.StatusOK, data)
	}
}
