package Service

import (
	"strconv"

	model "../Model"
	stash "../StashImportCycle"
)

func SetLoto7Predictions(param *stash.RegisterLoto7Predictions) bool {
	input_id, _ := strconv.Atoi(param.Body.UserID)
	input_data := make(map[int]map[string]int)

	for i := 0; i < len(param.Body.Predictions); i++ {
		input_data[i] = make(map[string]int)
		input_data[i]["time"], _ = strconv.Atoi(param.Body.Predictions[i]["time"])
		input_data[i]["input_number_1"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_1"])
		input_data[i]["input_number_2"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_2"])
		input_data[i]["input_number_3"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_3"])
		input_data[i]["input_number_4"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_4"])
		input_data[i]["input_number_5"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_5"])
		input_data[i]["input_number_6"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_6"])
		input_data[i]["input_number_7"], _ = strconv.Atoi(param.Body.Predictions[i]["input_number_7"])
	}

	dbResult := model.SetLoto7Predictions(input_data, input_id)
	return dbResult
}
