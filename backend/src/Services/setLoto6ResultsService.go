package Services

import (
	"strconv"

	model "../Models"
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

	var input_row [6]string
	input_row[0] = param["input_number_1"]
	input_row[1] = param["input_number_2"]
	input_row[2] = param["input_number_3"]
	input_row[3] = param["input_number_4"]
	input_row[4] = param["input_number_5"]
	input_row[5] = param["input_number_6"]

	dbResult := model.SetLoto6Results(input_data, input_row)
	return dbResult
}
