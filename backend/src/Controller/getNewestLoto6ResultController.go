package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

func GetNewestLoto6Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetNewestLoto6Result()
		return c.JSON(http.StatusOK, data)
	}
}
