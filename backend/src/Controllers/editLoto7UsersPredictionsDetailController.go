package Controllers

import (
	"net/http"

	service "../Services"
	stash "../StashImportCycle"
	"github.com/labstack/echo"
)

func EditLoto7UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		stashParam := &stash.EditLoto7PredictionsDetailStash{}
		if err := c.Bind(stashParam); err != nil {
			return err
		}
		result := service.EditLoto7UsersPredictionsDetail(stashParam)
		return c.JSON(http.StatusCreated, result)
	}
}
