package Service

import (
	model "../Model"
)

func GetNewestLoto7Result() []*model.Loto7NewestTime {
	data := model.GetNewestLoto7Result()
	data[0].Time = data[0].Time + 1
	return data
}
