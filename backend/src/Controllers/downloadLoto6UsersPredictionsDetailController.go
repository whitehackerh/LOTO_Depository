package Controllers

import (
	service "../Services"
	"github.com/labstack/echo"
)

type DownloadLoto6PredictionsDetailRequestParam struct {
	Body map[string]string `json:"body"`
}

func DownloadLoto6UsersPredictionsDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(DownloadLoto6PredictionsDetailRequestParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		service.DownloadLoto6UsersPredictionsDetail(param.Body)
		return c.File("Loto6PredictionsDetail.csv")
	}
}
