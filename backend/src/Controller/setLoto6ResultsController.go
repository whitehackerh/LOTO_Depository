package Controller

import (
	"fmt"
	"net/http"

	service "../Service"

	"github.com/labstack/echo"
)

type RegisterLoto6Results struct {
	// Time   string `json:"time"`
	// Data_1 string `json:"data_1"`
	// Data_2 string `json:"data_2"`
	// Data_3 string `json:"data_3"`
	// Data_4 string `json:"data_4"`
	// Data_5 string `json:"data_5"`
	// Data_6 string `json:"data_6"`
	Body map[string]string `json:"body"`
	// Time   int `json:"time"`
	// Data_1 int `json:"data_1"`
	// Data_2 int `json:"data_2"`
	// Data_3 int `json:"data_3"`
	// Data_4 int `json:"data_4"`
	// Data_5 int `json:"data_5"`
	// Data_6 int `json:"data_6"`
}

func SetLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		// setResults := new(RegisterLoto6Results)
		// a := c.FormValue("time")
		// b := c.FormValue("data_1")
		// fmt.Println(a)
		// fmt.Println(b)
		// a := c.FormValue("body")
		// b := c.Request().Body
		//	d := c.Param("time")
		//ex := new(RegisterLoto6Results)

		// if err := c.Bind(&ex); err != nil {
		// 	fmt.Println("error")
		// }
		// bi := c.Bind(ex)
		// fmt.Println(bi)
		// var abc RegisterLoto6Results
		// fmt.Println(c.Bind(abc))
		param := new(RegisterLoto6Results)
		// バインドしてJSON取得
		if err := c.Bind(param); err != nil {
			//return err
		}
		fmt.Println(param)
		fmt.Println(param.Body["data_1"])
		//fmt.Println(param.Time)
		result := service.SetLoto6Results(param.Body)
		fmt.Println(result)
		return c.JSON(http.StatusCreated, result)
	}
}
