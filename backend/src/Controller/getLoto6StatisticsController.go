package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

func GetLoto6Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto6Statistics()
		return c.JSON(http.StatusOK, data)
	}
}
