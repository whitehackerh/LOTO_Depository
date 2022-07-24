package Service

import (
	"strconv"

	model "../Model"
)

func GetLoto7ResultDetail(body map[string]string) []*model.Loto7ResultDetail {
	param, _ := strconv.Atoi(body["time"])
	data := model.GetLoto7ResultDetail(param)
	return data
}
