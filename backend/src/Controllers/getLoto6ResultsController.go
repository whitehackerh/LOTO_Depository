package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

func GetLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto6Results()
		return c.JSON(http.StatusOK, data)
	}
}
