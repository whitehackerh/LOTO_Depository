package Services

import (
	"strconv"

	model "../Models"
)

func GetLoto7ResultDetail(body map[string]string) []*model.Loto7ResultDetail {
	param, _ := strconv.Atoi(body["time"])
	data := model.GetLoto7ResultDetail(param)
	return data
}
