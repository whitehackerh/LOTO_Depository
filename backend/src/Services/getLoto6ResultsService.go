package Services

import (
	model "loto_depository/src/Models"
)

func GetLoto6Results() []*model.Loto6Results {
	data := model.GetLoto6Results()
	return data
}
