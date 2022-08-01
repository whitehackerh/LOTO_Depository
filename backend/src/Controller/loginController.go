package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type LoginInfo struct {
	Body map[string]string `json:"body"`
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(LoginInfo)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.Login(param.Body)
		return c.JSONBlob(http.StatusOK, data)
	}
}
