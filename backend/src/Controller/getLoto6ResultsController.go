package Controller

import (
	"net/http"

	service "../Service"
	"github.com/labstack/echo"
)

// type Results struct {
// 	Time             int `json:"time"`
// 	Winning_number_1 int `json:"winning_number_1"`
// 	Winning_number_2 int `json:"winning_number_2"`
// 	Winning_number_3 int `json:"winning_number_3"`
// 	Winning_number_4 int `json:"winning_number_4"`
// 	Winning_number_5 int `json:"winning_number_5"`
// 	Winning_number_6 int `json:"winning_number_6"`
// }

// func GetLoto6Results() echo.HandlerFunc {
// 	json := service.GetLoto6Results()
// 	return json
// }

func GetLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto6Results()
		return c.JSON(http.StatusOK, data)
	}
}

// func GetLoto6Results() []byte {
// 	json := service.GetLoto6Results()
// 	return json
// }
