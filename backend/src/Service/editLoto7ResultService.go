package Service

import (
	"strconv"

	model "../Model"
)

func EditLoto7Result(body map[string]string) bool {
	param := body
	input_data := make(map[string]int)
	time, _ := strconv.Atoi(param["time"])
	input_data["number_1"], _ = strconv.Atoi(param["input_number_1"])
	input_data["number_2"], _ = strconv.Atoi(param["input_number_2"])
	input_data["number_3"], _ = strconv.Atoi(param["input_number_3"])
	input_data["number_4"], _ = strconv.Atoi(param["input_number_4"])
	input_data["number_5"], _ = strconv.Atoi(param["input_number_5"])
	input_data["number_6"], _ = strconv.Atoi(param["input_number_6"])
	input_data["number_7"], _ = strconv.Atoi(param["input_number_7"])

	before_edit_record := model.GetLoto7ResultDetail(time)
	before_edit := make(map[string]int)
	before_edit["number_1"] = before_edit_record[0].Number_1
	before_edit["number_2"] = before_edit_record[0].Number_2
	before_edit["number_3"] = before_edit_record[0].Number_3
	before_edit["number_4"] = before_edit_record[0].Number_4
	before_edit["number_5"] = before_edit_record[0].Number_5
	before_edit["number_6"] = before_edit_record[0].Number_6
	before_edit["number_7"] = before_edit_record[0].Number_7

	// input / before 両方に存在 → statistics updateなし
	// input にのみ存在 → statistics count + 1, rate = (count + 1) / time
	// before にのみ存在 → statistics count - 1, rate = (count - 1) / time
	exist_only_input := false
	exist_only_before := false
	updateStatisticsQuery := make(map[int]string)
	params := make(map[int]int)
	counter := 0
	for _, number := range input_data {
		for _, before_number := range before_edit {
			if number == before_number {
				exist_only_input = true
				break
			}
		}
		if exist_only_input == false {
			updateStatisticsQuery[counter] = `UPDATE loto7_statistics SET count = (SELECT count + 1 FROM loto7_statistics WHERE number = $1), rate = (SELECT CAST(count + 1 AS float) FROM loto7_statistics WHERE number = $2) / (SELECT time FROM loto7_statistics WHERE number = $3) WHERE number = $4`
			params[counter] = number
			counter++
		}
		exist_only_input = false
	}
	for _, before_number := range before_edit {
		for _, number := range input_data {
			if before_number == number {
				exist_only_before = true
				break
			}
		}
		if exist_only_before == false {
			updateStatisticsQuery[counter] = `UPDATE loto7_statistics SET count = (SELECT count - 1 FROM loto7_statistics WHERE number = $1), rate = (SELECT CAST(count - 1 AS float) FROM loto7_statistics WHERE number = $2) / (SELECT time FROM loto7_statistics WHERE number = $3) WHERE number = $4`
			params[counter] = before_number
			counter++
		}
		exist_only_before = false
	}

	updateLoto7ResultsQuery := `UPDATE loto7_results SET number_1 = $1, number_2 = $2, number_3 = $3, number_4 = $4, number_5 = $5, number_6 = $6, number_7 = $7 WHERE time = $8`

	dbResult := model.UpdateLoto7Results(updateLoto7ResultsQuery, input_data, updateStatisticsQuery, params, time)
	return dbResult
}
