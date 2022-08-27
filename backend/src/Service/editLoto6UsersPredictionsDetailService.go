package Service

import (
	"strconv"

	model "../Model"
	stash "../StashImportCycle"
)

func EditLoto6UsersPredictionsDetail(param *stash.EditLoto6PredictionsDetailStash) bool {
	input_id, _ := strconv.Atoi(param.Body.UserID)
	input_data := make(map[int]map[string]int)

	for i := 0; i < len(param.Body.Predictions); i++ {
		input_data[i] = make(map[string]int)
		input_data[i]["time"], _ = strconv.Atoi(param.Body.Predictions[i]["Time"])
		input_data[i]["time_id"], _ = strconv.Atoi(param.Body.Predictions[i]["Time_Id"])
		input_data[i]["input_number_1"], _ = strconv.Atoi(param.Body.Predictions[i]["Number_1"])
		input_data[i]["input_number_2"], _ = strconv.Atoi(param.Body.Predictions[i]["Number_2"])
		input_data[i]["input_number_3"], _ = strconv.Atoi(param.Body.Predictions[i]["Number_3"])
		input_data[i]["input_number_4"], _ = strconv.Atoi(param.Body.Predictions[i]["Number_4"])
		input_data[i]["input_number_5"], _ = strconv.Atoi(param.Body.Predictions[i]["Number_5"])
		input_data[i]["input_number_6"], _ = strconv.Atoi(param.Body.Predictions[i]["Number_6"])
	}

	dbResult := model.EditLoto6UsersPredictionsDetail(input_data, input_id)
	return dbResult
}
