package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

func GetNewestLoto7Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetNewestLoto7Result()
		return c.JSON(http.StatusOK, data)
	}
}
