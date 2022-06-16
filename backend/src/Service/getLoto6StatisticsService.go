package Service

import (
	model "../Model"
)

func GetLoto6Statistics() []*model.Loto6Statistics {
	data := model.GetLoto6Statistics()
	for i := 0; i < len(data); i++ {
		data[i].Rate = data[i].Rate * 100
	}
	return data
}
