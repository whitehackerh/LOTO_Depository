package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	stash "loto_depository/src/StashImportCycle"

	"github.com/labstack/echo/v4"
)

func EditLoto6UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.EditLoto6PredictionsDetailStash{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.EditLoto6UsersPredictionsDetail(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
