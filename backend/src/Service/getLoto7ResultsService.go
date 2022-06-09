package Service

import (
	model "../Model"
)

func GetLoto7Results() []*model.Loto7Results {
	data := model.GetLoto7Results()
	return data
}
