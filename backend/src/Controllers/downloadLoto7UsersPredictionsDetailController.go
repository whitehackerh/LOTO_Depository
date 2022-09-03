package Controllers

import (
	service "../Services"
	"github.com/labstack/echo"
)

type DownloadLoto7PredictionsDetailRequestParam struct {
	Body map[string]string `json:"body"`
}

func DownloadLoto7UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(DownloadLoto7PredictionsDetailRequestParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		service.DownloadLoto7UsersPredictionsDetail(param.Body)
		return c.File("Loto7PredictionsDetail.csv")
	}
}
