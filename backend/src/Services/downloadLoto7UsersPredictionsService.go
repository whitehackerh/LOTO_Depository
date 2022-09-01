package Services

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	model "../Models"
)

type Loto7UsersPredictionsForCsv struct {
	BeforeResultAnnouncementPredictions map[int]map[string]int
	Records                             int
	Average                             float64
	Rate                                float64
	Predictions                         map[int]map[string]int
}

func DownloadLoto7UsersPredictions(body map[string]string) {
	user_id, _ := strconv.Atoi(body["user_id"])
	records := model.GetLoto7UsersPredictions(user_id)
	results := model.GetLoto7Results()
	latestResult := model.GetNewestLoto7Result()

	param, existBeforeResultAnnouncement, existAfterResultAnnouncement := makeParamLoto7(records, results, latestResult)
	sliceParamBefore, sliceParamAfter := castSliceParamLoto7(param, existBeforeResultAnnouncement, existAfterResultAnnouncement)
	if existBeforeResultAnnouncement == true {
		sort.Slice(sliceParamBefore, func(i, j int) bool {
			// Time DESC
			if sliceParamBefore[i][0].(int) > sliceParamBefore[j][0].(int) {
				return true
			} else if sliceParamBefore[i][0].(int) == sliceParamBefore[j][0].(int) {
				// Time_Id ASC
				if sliceParamBefore[i][1].(int) < sliceParamBefore[j][1].(int) {
					return true
				}
			}
			return false
		})
	}
	if existAfterResultAnnouncement == true {
		sort.Slice(sliceParamAfter, func(i, j int) bool {
			if sliceParamAfter[i][0].(int) > sliceParamAfter[j][0].(int) {
				return true
			} else if sliceParamAfter[i][0].(int) == sliceParamAfter[j][0].(int) {
				if sliceParamAfter[i][1].(int) < sliceParamAfter[j][1].(int) {
					return true
				}
			}
			return false
		})
	}
	castedParam := castParamLoto7(param, sliceParamBefore, sliceParamAfter, existBeforeResultAnnouncement, existAfterResultAnnouncement)

	f, err := os.Create("Loto7Predictions.csv")
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(castedParam)
}

func makeParamLoto7(records []*model.Loto7UsersPredictions, results []*model.Loto7Results, latestResult []*model.Loto7NewestTime) (Loto7UsersPredictionsForCsv, bool, bool) {
	param := Loto7UsersPredictionsForCsv{}

	beforeCounter := 0
	afterCounter := 0
	totalMatch := 0
	param.BeforeResultAnnouncementPredictions = make(map[int]map[string]int)
	param.Predictions = make(map[int]map[string]int)

	for _, value := range records {
		if value.Time > latestResult[0].Time {
			param.BeforeResultAnnouncementPredictions[beforeCounter] = make(map[string]int)
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Time"] = value.Time
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Time_Id"] = value.Time_Id
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_1"] = value.Number_1
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_2"] = value.Number_2
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_3"] = value.Number_3
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_4"] = value.Number_4
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_5"] = value.Number_5
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_6"] = value.Number_6
			param.BeforeResultAnnouncementPredictions[beforeCounter]["Number_7"] = value.Number_7
			beforeCounter++
		} else {
			param.Predictions[afterCounter] = make(map[string]int)
			param.Predictions[afterCounter]["Time"] = value.Time
			param.Predictions[afterCounter]["Time_Id"] = value.Time_Id
			param.Predictions[afterCounter]["Number_1"] = value.Number_1
			param.Predictions[afterCounter]["Number_2"] = value.Number_2
			param.Predictions[afterCounter]["Number_3"] = value.Number_3
			param.Predictions[afterCounter]["Number_4"] = value.Number_4
			param.Predictions[afterCounter]["Number_5"] = value.Number_5
			param.Predictions[afterCounter]["Number_6"] = value.Number_6
			param.Predictions[afterCounter]["Number_7"] = value.Number_7
			afterCounter++
		}
	}

	if beforeCounter == 0 && afterCounter == 0 {
		return param, false, false
	} else if beforeCounter != 0 && afterCounter == 0 {
		return param, true, false
	}

	for i := 0; i < len(param.Predictions); i++ {
		param.Predictions[i]["Matches"] = 0
		for _, value := range results {
			if param.Predictions[i]["Time"] == value.Time {
				if param.Predictions[i]["Number_1"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_1"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_1"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_1"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_1"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_1"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_1"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_2"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_3"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_4"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_5"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_6"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_1 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_2 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_3 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_4 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_5 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_6 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
				if param.Predictions[i]["Number_7"] == value.Number_7 {
					param.Predictions[i]["Matches"]++
					totalMatch++
				}
			}
		}
	}
	param.Records = len(param.Predictions)
	shift := math.Pow(10, 3)
	shift2 := math.Pow(10, 5)
	param.Average = math.Trunc((float64(totalMatch)/float64(param.Records))*shift) / shift
	param.Rate = math.Trunc((float64(totalMatch)/(float64(param.Records)*float64(7)))*shift2) / shift2 * float64(100)
	if beforeCounter == 0 {
		return param, false, true
	} else {
		return param, true, true
	}
}

func castSliceParamLoto7(param Loto7UsersPredictionsForCsv, existBeforeResultAnnouncement bool, existAfterResultAnnouncement bool) ([][]interface{}, [][]interface{}) {
	sliceParamBefore := make([][]interface{}, 0, 0)
	sliceParamAfter := make([][]interface{}, 0, 0)
	if existBeforeResultAnnouncement == true {
		for _, value := range param.BeforeResultAnnouncementPredictions {
			sliceParamBefore = append(sliceParamBefore, []interface{}{value["Time"], value["Time_Id"], value["Number_1"], value["Number_2"], value["Number_3"], value["Number_4"], value["Number_5"], value["Number_6"], value["Number_7"]})
		}
	}
	if existAfterResultAnnouncement == true {
		for _, value := range param.Predictions {
			sliceParamAfter = append(sliceParamAfter, []interface{}{value["Time"], value["Time_Id"], value["Number_1"], value["Number_2"], value["Number_3"], value["Number_4"], value["Number_5"], value["Number_6"], value["Number_7"], value["Matches"]})
		}
	}
	return sliceParamBefore, sliceParamAfter
}

func castParamLoto7(param Loto7UsersPredictionsForCsv, sliceParamBefore [][]interface{}, sliceParamAfter [][]interface{}, existBeforeResultAnnouncement bool, existAfterResultAnnouncement bool) [][]string {
	castedParam := make([][]string, 0, 0)

	if existBeforeResultAnnouncement == true {
		// title
		castedParamBeforeTitle := make([]string, 1)
		castedParamBeforeTitle[0] = "Predictions of before result announcement"
		castedParam = append(castedParam, castedParamBeforeTitle)

		// header
		castedParamBeforeHeader := make([]string, 8)
		castedParamBeforeHeader[0] = "Time"
		castedParamBeforeHeader[1] = "Number1"
		castedParamBeforeHeader[2] = "Number2"
		castedParamBeforeHeader[3] = "Number3"
		castedParamBeforeHeader[4] = "Number4"
		castedParamBeforeHeader[5] = "Number5"
		castedParamBeforeHeader[6] = "Number6"
		castedParamBeforeHeader[7] = "Number7"
		castedParam = append(castedParam, castedParamBeforeHeader)

		// records
		for i := 0; i < len(sliceParamBefore); i++ {
			castedParamBeforeRecords := make([]string, 8)
			castedParamBeforeRecords[0] = strconv.Itoa(sliceParamBefore[i][0].(int))
			castedParamBeforeRecords[1] = strconv.Itoa(sliceParamBefore[i][2].(int))
			castedParamBeforeRecords[2] = strconv.Itoa(sliceParamBefore[i][3].(int))
			castedParamBeforeRecords[3] = strconv.Itoa(sliceParamBefore[i][4].(int))
			castedParamBeforeRecords[4] = strconv.Itoa(sliceParamBefore[i][5].(int))
			castedParamBeforeRecords[5] = strconv.Itoa(sliceParamBefore[i][6].(int))
			castedParamBeforeRecords[6] = strconv.Itoa(sliceParamBefore[i][7].(int))
			castedParamBeforeRecords[7] = strconv.Itoa(sliceParamBefore[i][8].(int))
			castedParam = append(castedParam, castedParamBeforeRecords)
		}

		// blank
		blank := make([]string, 1)
		castedParam = append(castedParam, blank)
	}

	if existAfterResultAnnouncement == true {
		// title
		castedParamAfterTitle := make([]string, 1)
		castedParamAfterTitle[0] = "Predictions of after result announcement"
		castedParam = append(castedParam, castedParamAfterTitle)

		// header
		castedParamAfterHeaderStatistics := make([]string, 3)
		castedParamAfterHeaderStatistics[0] = "Records"
		castedParamAfterHeaderStatistics[1] = "Average"
		castedParamAfterHeaderStatistics[2] = "Rate"
		castedParam = append(castedParam, castedParamAfterHeaderStatistics)

		// data
		castedParamAfterStatistics := make([]string, 3)
		castedParamAfterStatistics[0] = strconv.Itoa(param.Records)
		castedParamAfterStatistics[1] = strconv.FormatFloat(param.Average, 'f', -1, 64)
		castedParamAfterStatistics[2] = strconv.FormatFloat(param.Rate, 'f', -1, 64) + "%"
		castedParam = append(castedParam, castedParamAfterStatistics)

		// header
		castedParamAfterHeader := make([]string, 9)
		castedParamAfterHeader[0] = "Time"
		castedParamAfterHeader[1] = "Number1"
		castedParamAfterHeader[2] = "Number2"
		castedParamAfterHeader[3] = "Number3"
		castedParamAfterHeader[4] = "Number4"
		castedParamAfterHeader[5] = "Number5"
		castedParamAfterHeader[6] = "Number6"
		castedParamAfterHeader[7] = "Number7"
		castedParamAfterHeader[8] = "Matches"
		castedParam = append(castedParam, castedParamAfterHeader)

		// records
		for i := 0; i < len(sliceParamAfter); i++ {
			castedParamAfterRecords := make([]string, 9)
			castedParamAfterRecords[0] = strconv.Itoa(sliceParamAfter[i][0].(int))
			castedParamAfterRecords[1] = strconv.Itoa(sliceParamAfter[i][2].(int))
			castedParamAfterRecords[2] = strconv.Itoa(sliceParamAfter[i][3].(int))
			castedParamAfterRecords[3] = strconv.Itoa(sliceParamAfter[i][4].(int))
			castedParamAfterRecords[4] = strconv.Itoa(sliceParamAfter[i][5].(int))
			castedParamAfterRecords[5] = strconv.Itoa(sliceParamAfter[i][6].(int))
			castedParamAfterRecords[6] = strconv.Itoa(sliceParamAfter[i][7].(int))
			castedParamAfterRecords[7] = strconv.Itoa(sliceParamAfter[i][8].(int))
			castedParamAfterRecords[8] = strconv.Itoa(sliceParamAfter[i][9].(int))
			castedParam = append(castedParam, castedParamAfterRecords)
		}
	}
	return castedParam
}
