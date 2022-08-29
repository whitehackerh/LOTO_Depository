package Services

import (
	"strconv"

	model "../Models"
)

func GetLoto6ResultDetail(body map[string]string) []*model.Loto6ResultDetail {
	param, _ := strconv.Atoi(body["time"])
	data := model.GetLoto6ResultDetail(param)
	return data
}
