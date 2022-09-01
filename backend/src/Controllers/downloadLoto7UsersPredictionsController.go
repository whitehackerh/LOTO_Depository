package Controllers

import (
	service "../Services"
	"github.com/labstack/echo"
)

type Loto7PredictionsUserIdForCsv struct {
	Body map[string]string `json:"body"`
}

func DownloadLoto7UsersPredictions() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto7PredictionsUserIdForCsv)
		if err := c.Bind(param); err != nil {
			return err
		}
		service.DownloadLoto7UsersPredictions(param.Body)
		return c.File("Loto7Predictions.csv")
	}
}
