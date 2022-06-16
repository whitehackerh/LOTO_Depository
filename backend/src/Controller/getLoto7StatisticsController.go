package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

func GetLoto7Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto7Statistics()
		return c.JSON(http.StatusOK, data)
	}
}
