package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type DetermineLoto6Param struct {
	Body map[string]string `json:"body"`
}

func DetermineLoto6Expectation() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(DetermineLoto6Param)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.DetermineLoto6Expectation(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
