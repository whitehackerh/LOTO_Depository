package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

type CreateUserForm struct {
	Body map[string]string `json:"body"`
}

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(CreateUserForm)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.CreateUser(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
