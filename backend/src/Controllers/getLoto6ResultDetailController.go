package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

type Loto6ResultDetailParam struct {
	Body map[string]string `json:"body"`
}

func GetLoto6ResultDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto6ResultDetailParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.GetLoto6ResultDetail(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
