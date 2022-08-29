package Services

import (
	model "../Models"
)

func GetLoto7Results() []*model.Loto7Results {
	data := model.GetLoto7Results()
	return data
}
