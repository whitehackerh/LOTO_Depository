package Services

import (
	"strconv"

	model "../Models"
)

func GetLoto6UsersPredictionsDetail(body map[string]string) map[string]map[int]map[string]interface{} {
	user_id, _ := strconv.Atoi(body["user_id"])
	time, _ := strconv.Atoi(body["time"])
	records := model.GetLoto6UsersPredictionsDetail(user_id, time)
	latestResult := model.GetNewestLoto6Result()

	data := make(map[string]map[int]map[string]interface{})
	data["Result"] = make(map[int]map[string]interface{})
	data["Result"][0] = make(map[string]interface{})
	results := []*model.Loto6ResultDetail{}

	beforeResultAnnouncement := false

	if time <= latestResult[0].Time {
		results = model.GetLoto6ResultDetail(time)
	} else {
		results = append(results, &model.Loto6ResultDetail{Time: 0, Number_1: 0, Number_2: 0, Number_3: 0, Number_4: 0, Number_5: 0, Number_6: 0})
		beforeResultAnnouncement = true
	}

	/*
				case 1 結果発表前 -> 予想のみ表示
		        case 2 結果発表後 -> 結果・予想両方表示
				case 3 予想がない -> 結果のみ表示
				case 4 結果も予想もない -> どちらも非表示
	*/

	data["Result"][0]["Time"] = results[0].Time
	data["Result"][0]["Number_1"] = results[0].Number_1
	data["Result"][0]["Number_2"] = results[0].Number_2
	data["Result"][0]["Number_3"] = results[0].Number_3
	data["Result"][0]["Number_4"] = results[0].Number_4
	data["Result"][0]["Number_5"] = results[0].Number_5
	data["Result"][0]["Number_6"] = results[0].Number_6

	data["Predictions"] = make(map[int]map[string]interface{})
	if records[0].Time == 0 {
		data["Predictions"][0]["Records"] = 0
		data["Predictions"][0]["Total"] = 0
		data["Predictions"][0]["Rate"] = 0.0
		data["Predictions"][0]["Average"] = 0.0
		data["Predictions"][0]["Time"] = 0
		data["Predictions"][0]["beforeResultAnnouncement"] = beforeResultAnnouncement
		return data
	}

	for index, value := range records {
		data["Predictions"][index] = make(map[string]interface{})
		data["Predictions"][index]["User_id"] = value.User_Id
		data["Predictions"][index]["Time"] = value.Time
		data["Predictions"][index]["Time_Id"] = value.Time_Id
		data["Predictions"][index]["Number_1"] = value.Number_1
		data["Predictions"][index]["Number_2"] = value.Number_2
		data["Predictions"][index]["Number_3"] = value.Number_3
		data["Predictions"][index]["Number_4"] = value.Number_4
		data["Predictions"][index]["Number_5"] = value.Number_5
		data["Predictions"][index]["Number_6"] = value.Number_6
	}

	if beforeResultAnnouncement == true {
		data["Predictions"][0]["Total"] = 0
		data["Predictions"][0]["Records"] = len(data["Predictions"])
		data["Predictions"][0]["Average"] = 0
		data["Predictions"][0]["Rate"] = 0
		data["Predictions"][0]["beforeResultAnnouncement"] = true
		return data
	}

	total := 0
	for i := 0; i < len(data["Predictions"]); i++ {
		data["Predictions"][i]["Matches"] = 0
		if data["Predictions"][i]["Number_1"] == data["Result"][0]["Number_1"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_1"] == data["Result"][0]["Number_2"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_1"] == data["Result"][0]["Number_3"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_1"] == data["Result"][0]["Number_4"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_1"] == data["Result"][0]["Number_5"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_1"] == data["Result"][0]["Number_6"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_2"] == data["Result"][0]["Number_1"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_2"] == data["Result"][0]["Number_2"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_2"] == data["Result"][0]["Number_3"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_2"] == data["Result"][0]["Number_4"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_2"] == data["Result"][0]["Number_5"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_2"] == data["Result"][0]["Number_6"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_3"] == data["Result"][0]["Number_1"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_3"] == data["Result"][0]["Number_2"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_3"] == data["Result"][0]["Number_3"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_3"] == data["Result"][0]["Number_4"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_3"] == data["Result"][0]["Number_5"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_3"] == data["Result"][0]["Number_6"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_4"] == data["Result"][0]["Number_1"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_4"] == data["Result"][0]["Number_2"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_4"] == data["Result"][0]["Number_3"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_4"] == data["Result"][0]["Number_4"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_4"] == data["Result"][0]["Number_5"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_4"] == data["Result"][0]["Number_6"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_5"] == data["Result"][0]["Number_1"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_5"] == data["Result"][0]["Number_2"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_5"] == data["Result"][0]["Number_3"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_5"] == data["Result"][0]["Number_4"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_5"] == data["Result"][0]["Number_5"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_5"] == data["Result"][0]["Number_6"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_6"] == data["Result"][0]["Number_1"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_6"] == data["Result"][0]["Number_2"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_6"] == data["Result"][0]["Number_3"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_6"] == data["Result"][0]["Number_4"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_6"] == data["Result"][0]["Number_5"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
		if data["Predictions"][i]["Number_6"] == data["Result"][0]["Number_6"] {
			data["Predictions"][i]["Matches"] = data["Predictions"][i]["Matches"].(int) + 1
			total++
		}
	}
	data["Predictions"][0]["Total"] = total
	data["Predictions"][0]["Records"] = len(data["Predictions"])
	data["Predictions"][0]["Average"] = float64(total) / float64(len(data["Predictions"]))
	data["Predictions"][0]["Rate"] = float64(total) / (float64(len(data["Predictions"])) * 6) * 100
	data["Predictions"][0]["beforeResultAnnouncement"] = false
	return data
}
