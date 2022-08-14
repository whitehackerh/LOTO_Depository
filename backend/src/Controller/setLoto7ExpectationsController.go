package Controller

import (
	"net/http"

	service "../Service"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
)

func SetLoto7Expectations() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.RegisterLoto7Expectations{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.SetLoto7Expectations(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
