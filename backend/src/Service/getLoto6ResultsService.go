package Service

import (
	model "../Model"
)

func GetLoto6Results() []*model.Loto6Results {
	data := model.GetLoto6Results()
	// data[0].Time = 1     上書きの方法
	return data
}
