package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type Loto7PredictionsDetail struct {
	Body map[string]string `json:"body"`
}

func GetLoto7UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto7PredictionsDetail)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.GetLoto7UsersPredictionsDetail(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
