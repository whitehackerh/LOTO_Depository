package Services

import (
	model "loto_depository/src/Models"
)

func GetLoto7Results() []*model.Loto7Results {
	data := model.GetLoto7Results()
	return data
}
