package Controllers

import (
	"net/http"

	service "../Services"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
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
