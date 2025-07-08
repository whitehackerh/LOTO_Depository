package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

type EditLoto7ResultParam struct {
	Body map[string]string `json:"body"`
}

func EditLoto7Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(EditLoto7ResultParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		result := service.EditLoto7Result(param.Body)
		return c.JSON(http.StatusCreated, result)
	}
}
