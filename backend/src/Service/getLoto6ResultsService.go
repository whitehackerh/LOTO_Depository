package Service

import (
	model "../Model"
)

// func GetLoto6Results() echo.HandlerFunc {
// 	results := model.GetLoto6Results()
// 	return results
// }

func GetLoto6Results() []*model.Loto6Results {
	data := model.GetLoto6Results()
	// fmt.Println(data)
	// data[0].Time = 1     上書きの方法
	return data
}
