package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	stash "loto_depository/src/StashImportCycle"

	"github.com/labstack/echo/v4"
)

func DeleteLoto7UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.DeleteLoto7PredictionsDetailStash{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.DeleteLoto7UsersPredictionsDetail(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
