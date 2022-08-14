package Service

import (
	"strconv"

	model "../Model"
	stash "../StashImportCycle"
)

func SetLoto6Expectations(param *stash.RegisterLoto6Expectations) bool {
	input_id, _ := strconv.Atoi(param.Body.UserID)
	input_data := make(map[int]map[string]int)

	for i := 0; i < len(param.Body.Expectations); i++ {
		input_data[i] = make(map[string]int)
		input_data[i]["time"], _ = strconv.Atoi(param.Body.Expectations[i]["time"])
		input_data[i]["input_number_1"], _ = strconv.Atoi(param.Body.Expectations[i]["input_number_1"])
		input_data[i]["input_number_2"], _ = strconv.Atoi(param.Body.Expectations[i]["input_number_2"])
		input_data[i]["input_number_3"], _ = strconv.Atoi(param.Body.Expectations[i]["input_number_3"])
		input_data[i]["input_number_4"], _ = strconv.Atoi(param.Body.Expectations[i]["input_number_4"])
		input_data[i]["input_number_5"], _ = strconv.Atoi(param.Body.Expectations[i]["input_number_5"])
		input_data[i]["input_number_6"], _ = strconv.Atoi(param.Body.Expectations[i]["input_number_6"])
	}

	dbResult := model.SetLoto6Expectations(input_data, input_id)
	return dbResult
}
