package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
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
