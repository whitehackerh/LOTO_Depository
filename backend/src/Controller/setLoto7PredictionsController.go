package Controller

import (
	"net/http"

	service "../Service"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
)

func SetLoto7Predictions() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.RegisterLoto7Predictions{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.SetLoto7Predictions(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
