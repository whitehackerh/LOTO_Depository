package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type RegisterLoto7Results struct {
	Body map[string]string `json:"body"`
}

func SetLoto7Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(RegisterLoto7Results)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.SetLoto7Results(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
