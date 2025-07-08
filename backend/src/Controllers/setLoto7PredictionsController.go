package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	stash "loto_depository/src/StashImportCycle"

	"github.com/labstack/echo/v4"
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
