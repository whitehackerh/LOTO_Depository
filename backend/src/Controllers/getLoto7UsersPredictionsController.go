package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

type Loto7PredictionsUserId struct {
	Body map[string]string `json:"body"`
}

func GetLoto7UsersPredictions() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto7PredictionsUserId)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.GetLoto7UsersPredictions(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
