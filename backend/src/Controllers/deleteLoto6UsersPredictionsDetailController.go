package Controllers

import (
	"net/http"

	service "../Services"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
)

func DeleteLoto6UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.DeleteLoto6PredictionsDetailStash{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.DeleteLoto6UsersPredictionsDetail(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
