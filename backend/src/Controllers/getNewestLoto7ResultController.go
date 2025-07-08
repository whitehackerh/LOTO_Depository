package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func GetNewestLoto7Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetNewestLoto7Result()
		return c.JSON(http.StatusOK, data)
	}
}
