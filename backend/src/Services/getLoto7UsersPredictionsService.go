package Services

import (
	"strconv"

	model "../Models"
)

func GetLoto7UsersPredictions(body map[string]string) map[int]map[string]interface{} {
	user_id, _ := strconv.Atoi(body["user_id"])
	records := model.GetLoto7UsersPredictions(user_id)
	results := model.GetLoto7Results()
	latestResult := model.GetNewestLoto7Result()
	beforeResultAnnouncementCount := 0

	data := make(map[int]map[string]interface{})
	if records[0].Time == 0 {
		data[0] = make(map[string]interface{})
		data[0]["Records"] = 0
		data[0]["Total"] = 0
		data[0]["Rate"] = 0.0
		data[0]["Average"] = 0.0
		data[0]["Time"] = 0
		data[0]["existBeforeResultAnnouncement"] = false
		data[0]["existAfterResultAnnouncement"] = false
		data[0]["beforeResultAnnouncementCount"] = 0
		return data
	}

	counter := 0
	for _, value := range records {
		data[counter] = make(map[string]interface{})
		data[counter]["Time"] = value.Time
		data[counter]["Number_1"] = value.Number_1
		data[counter]["Number_2"] = value.Number_2
		data[counter]["Number_3"] = value.Number_3
		data[counter]["Number_4"] = value.Number_4
		data[counter]["Number_5"] = value.Number_5
		data[counter]["Number_6"] = value.Number_6
		data[counter]["Number_7"] = value.Number_7
		counter++
	}

	counter2 := 0
	total := 0
	data[0]["existAfterResultAnnouncement"] = false
	data[0]["existAfterResultAnnouncement"] = false
	for i := 0; i < len(data); i++ {
		data[counter2]["Matches"] = 0
		data[counter2]["beforeResultAnnouncement"] = false
		if data[counter2]["Time"].(int) > latestResult[0].Time {
			beforeResultAnnouncementCount++
			data[counter2]["beforeResultAnnouncement"] = true
			counter2++
			continue
		}
		data[0]["existAfterResultAnnouncement"] = true
		for _, value := range results {
			if data[counter2]["Time"] == value.Time {
				if data[counter2]["Number_1"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_1"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_1"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_1"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_1"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_1"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_1"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_2"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_3"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_4"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_5"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_6"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_1 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_2 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_3 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_4 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_5 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_6 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
				if data[counter2]["Number_7"] == value.Number_7 {
					data[counter2]["Matches"] = data[counter2]["Matches"].(int) + 1
					total++
				}
			}
		}
		counter2++
	}
	data[0]["beforeResultAnnouncementCount"] = beforeResultAnnouncementCount
	if beforeResultAnnouncementCount > 0 {
		data[0]["existBeforeResultAnnouncement"] = true
	} else {
		data[0]["existBeforeResultAnnouncement"] = false
	}
	data[0]["Total"] = total
	data[0]["Records"] = len(data) - beforeResultAnnouncementCount
	if data[0]["Records"].(int) <= 0 {
		data[0]["Average"] = 0.0
		data[0]["Rate"] = 0.0
		data[0]["existAfterResultAnnouncement"] = false
		return data
	}
	data[0]["Average"] = float64(total) / (float64(len(data)) - float64(beforeResultAnnouncementCount))
	data[0]["Rate"] = float64(total) / ((float64(len(data)) - float64(beforeResultAnnouncementCount)) * float64(7)) * float64(100)
	data[0]["existAfterResultAnnouncement"] = true
	return data
}
