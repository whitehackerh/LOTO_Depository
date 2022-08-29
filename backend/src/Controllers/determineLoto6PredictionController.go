package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

type DetermineLoto6Param struct {
	Body map[string]string `json:"body"`
}

func DetermineLoto6Prediction() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(DetermineLoto6Param)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.DetermineLoto6Prediction(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
