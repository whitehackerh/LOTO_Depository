package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	stash "loto_depository/src/StashImportCycle"

	"github.com/labstack/echo/v4"
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
