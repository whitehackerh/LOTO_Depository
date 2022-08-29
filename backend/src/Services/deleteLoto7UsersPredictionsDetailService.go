package Services

import (
	"strconv"

	model "../Models"
	stash "../StashImportCycle"
)

func DeleteLoto7UsersPredictionsDetail(param *stash.DeleteLoto7PredictionsDetailStash) bool {
	input_id, _ := strconv.Atoi(param.Body.UserID)
	delete_data := make(map[int]map[string]int)

	for i := 0; i < len(param.Body.Predictions); i++ {
		delete_data[i] = make(map[string]int)
		delete_data[i]["time"], _ = strconv.Atoi(param.Body.Predictions[i]["Time"])
		delete_data[i]["time_id"], _ = strconv.Atoi(param.Body.Predictions[i]["Time_Id"])
	}

	dbResult := model.DeleteLoto7UsersPredictionsDetail(delete_data, input_id)
	return dbResult
}
