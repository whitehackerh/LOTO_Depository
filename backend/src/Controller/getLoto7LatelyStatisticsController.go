package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

func GetLoto7LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto7LatelyStatistics()
		return c.JSON(http.StatusOK, data)
	}
}
