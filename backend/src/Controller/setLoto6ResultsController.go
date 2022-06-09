package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type RegisterLoto6Results struct {
	Body map[string]string `json:"body"`
}

func SetLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(RegisterLoto6Results)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.SetLoto6Results(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
