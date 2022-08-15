package Controller

import (
	"net/http"

	service "../Service"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
)

func SetLoto6Predictions() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.RegisterLoto6Predictions{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.SetLoto6Predictions(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
