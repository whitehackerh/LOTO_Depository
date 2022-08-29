package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

type Loto6PredictionsUserId struct {
	Body map[string]string `json:"body"`
}

func GetLoto6UsersPredictions() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto6PredictionsUserId)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.GetLoto6UsersPredictions(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
