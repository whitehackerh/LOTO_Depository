package Service

import (
	"strconv"

	model "../Model"
)

func SetLoto6Results(body map[string]string) bool {
	param := body
	data := make(map[string]int)
	data["time"], _ = strconv.Atoi(param["time"])
	data["data_1"], _ = strconv.Atoi(param["data_1"])
	data["data_2"], _ = strconv.Atoi(param["data_2"])
	data["data_3"], _ = strconv.Atoi(param["data_3"])
	data["data_4"], _ = strconv.Atoi(param["data_4"])
	data["data_5"], _ = strconv.Atoi(param["data_5"])
	data["data_6"], _ = strconv.Atoi(param["data_6"])
	insertResult := model.SetLoto6Results(data)
	return insertResult
}
