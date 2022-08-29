package Controllers

import (
	"net/http"

	service "../Services"
	"github.com/labstack/echo"
)

type Loto7ResultDetailParam struct {
	Body map[string]string `json:"body"`
}

func GetLoto7ResultDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(Loto7ResultDetailParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		data := service.GetLoto7ResultDetail(param.Body)
		return c.JSON(http.StatusOK, data)
	}
}
