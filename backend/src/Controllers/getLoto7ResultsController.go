package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

func GetLoto7Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto7Results()
		return c.JSON(http.StatusOK, data)
	}
}
