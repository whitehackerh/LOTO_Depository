package Service

import (
	model "../Model"
)

func GetLoto6Results() []*model.Loto6Results {
	data := model.GetLoto6Results()
	return data
}
