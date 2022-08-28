package Controller

import (
	"net/http"

	service "../Service"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
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
