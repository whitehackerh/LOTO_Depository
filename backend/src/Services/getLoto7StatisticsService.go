package Services

import (
	model "../Models"
)

func GetLoto7Statistics() []*model.Loto7Statistics {
	data := model.GetLoto7Statistics()
	for i := 0; i < len(data); i++ {
		data[i].Rate = data[i].Rate * 100
	}
	return data
}
