package Controllers

import (
	service "../Services"
	"github.com/labstack/echo"
)

type Loto6PredictionsUserIdForCsv struct {
	Body map[string]string `json:"body"`
}

func DownloadLoto6UsersPredictions() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto6PredictionsUserIdForCsv)
		if err := c.Bind(param); err != nil {
			return err
		}
		service.DownloadLoto6UsersPredictions(param.Body)
		return c.File("Loto6Predictions.csv")
	}
}
