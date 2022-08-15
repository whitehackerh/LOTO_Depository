package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type DetermineLoto7Param struct {
	Body map[string]string `json:"body"`
}

func DetermineLoto7Prediction() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(DetermineLoto7Param)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.DetermineLoto7Prediction(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
