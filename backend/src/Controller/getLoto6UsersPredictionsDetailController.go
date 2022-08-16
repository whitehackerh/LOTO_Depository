package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

type Loto6PredictionsDetail struct {
	Body map[string]string `json:"body"`
}

func GetLoto6UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto6PredictionsDetail)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.GetLoto6UsersPredictionsDetail(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
