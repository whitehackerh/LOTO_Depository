package Services

import (
	model "../Models"
)

func GetLoto6Results() []*model.Loto6Results {
	data := model.GetLoto6Results()
	return data
}
