package Service

import (
	"strconv"

	model "../Model"
)

func SetLoto6Results(body map[string]string) bool {
	param := body
	input_data := make(map[string]int)
	input_data["time"], _ = strconv.Atoi(param["time"])
	input_data["input_number_1"], _ = strconv.Atoi(param["input_number_1"])
	input_data["input_number_2"], _ = strconv.Atoi(param["input_number_2"])
	input_data["input_number_3"], _ = strconv.Atoi(param["input_number_3"])
	input_data["input_number_4"], _ = strconv.Atoi(param["input_number_4"])
	input_data["input_number_5"], _ = strconv.Atoi(param["input_number_5"])
	input_data["input_number_6"], _ = strconv.Atoi(param["input_number_6"])

	var input_column [6]string
	input_column[0] = "n" + param["input_number_1"]
	input_column[1] = "n" + param["input_number_2"]
	input_column[2] = "n" + param["input_number_3"]
	input_column[3] = "n" + param["input_number_4"]
	input_column[4] = "n" + param["input_number_5"]
	input_column[5] = "n" + param["input_number_6"]

	dbResult := model.SetLoto6Results(input_data, input_column)
	return dbResult
}
