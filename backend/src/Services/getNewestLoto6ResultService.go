package Services

import (
	model "../Models"
)

func GetNewestLoto6Result() []*model.Loto6NewestTime {
	data := model.GetNewestLoto6Result()
	data[0].Time = data[0].Time + 1
	return data
}
