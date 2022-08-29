package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

type EditLoto6ResultParam struct {
	Body map[string]string `json:"body"`
}

func EditLoto6Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(EditLoto6ResultParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.EditLoto6Result(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
