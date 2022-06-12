package Service

import (
	"strconv"

	model "../Model"
)

type DetermineLoto7Response struct {
	Result string
}

func DetermineLoto7Expectation(body map[string]string) []*DetermineLoto7Response {
	param := body
	input_data := make(map[string]int)
	input_data["time"], _ = strconv.Atoi(param["time"])
	input_data["input_number_1"], _ = strconv.Atoi(param["input_number_1"])
	input_data["input_number_2"], _ = strconv.Atoi(param["input_number_2"])
	input_data["input_number_3"], _ = strconv.Atoi(param["input_number_3"])
	input_data["input_number_4"], _ = strconv.Atoi(param["input_number_4"])
	input_data["input_number_5"], _ = strconv.Atoi(param["input_number_5"])
	input_data["input_number_6"], _ = strconv.Atoi(param["input_number_6"])
	input_data["input_number_7"], _ = strconv.Atoi(param["input_number_7"])

	records := model.GetLoto7Results()
	judge := false
	for _, value := range records {
		if input_data["input_number_1"] != value.Number_1 {
			continue
		}
		if input_data["input_number_2"] != value.Number_2 {
			continue
		}
		if input_data["input_number_3"] != value.Number_3 {
			continue
		}
		if input_data["input_number_4"] != value.Number_4 {
			continue
		}
		if input_data["input_number_5"] != value.Number_5 {
			continue
		}
		if input_data["input_number_6"] != value.Number_6 {
			continue
		}
		if input_data["input_number_7"] != value.Number_7 {
			continue
		} else {
			judge = true
		}
		if judge {
			break
		}
	}
	determineStruct := DetermineLoto7Response{}
	if judge {
		determineStruct.Result = "All Match"
	} else {
		determineStruct.Result = "Not Match"
	}
	response := []*DetermineLoto7Response{}
	response = append(response, &DetermineLoto7Response{Result: determineStruct.Result})
	return response
}
