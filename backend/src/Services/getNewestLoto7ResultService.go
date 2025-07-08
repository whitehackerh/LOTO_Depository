package Services

import (
	model "loto_depository/src/Models"
)

func GetNewestLoto7Result() []*model.Loto7NewestTime {
	data := model.GetNewestLoto7Result()
	data[0].Time = data[0].Time + 1
	return data
}
