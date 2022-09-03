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

type Loto7UsersPredictionsDetailForCsv struct {
	Result      map[string]int
	Records     int
	Average     float64
	Rate        float64
	Predictions map[int]map[string]int
}

func DownloadLoto7UsersPredictionsDetail(body map[string]string) {
	user_id, _ := strconv.Atoi(body["user_id"])
	time, _ := strconv.Atoi(body["time"])
	records := model.GetLoto7UsersPredictionsDetail(user_id, time)
	latestResult := model.GetNewestLoto7Result()

	param, beforeResultAnnouncement, existPredictions := makeParamLoto7Detail(time, latestResult, records)
	sliceResult, slicePredictions := castSliceParamLoto7Detail(param, beforeResultAnnouncement, existPredictions)

	if existPredictions == true {
		sort.Slice(slicePredictions, func(i, j int) bool {
			if slicePredictions[i][0].(int) < slicePredictions[j][0].(int) {
				return true
			}
			return false
		})
	}
	castedParam := castParamLoto7Detail(time, param, sliceResult, slicePredictions, beforeResultAnnouncement, existPredictions)

	f, err := os.Create("Loto7PredictionsDetail.csv")
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(castedParam)
}

func makeParamLoto7Detail(time int, latestResult []*model.Loto7NewestTime, records []*model.Loto7UsersPredictions) (Loto7UsersPredictionsDetailForCsv, bool, bool) {
	beforeResultAnnouncement := false
	existPredictions := true
	param := Loto7UsersPredictionsDetailForCsv{}
	param.Result = make(map[string]int)

	if time <= latestResult[0].Time {
		result := model.GetLoto7ResultDetail(time)
		param.Result["Number_1"] = result[0].Number_1
		param.Result["Number_2"] = result[0].Number_2
		param.Result["Number_3"] = result[0].Number_3
		param.Result["Number_4"] = result[0].Number_4
		param.Result["Number_5"] = result[0].Number_5
		param.Result["Number_6"] = result[0].Number_6
		param.Result["Number_7"] = result[0].Number_7
	} else {
		beforeResultAnnouncement = true
	}

	param.Predictions = make(map[int]map[string]int)

	if records[0].Time == 0 {
		existPredictions = false
		return param, beforeResultAnnouncement, existPredictions
	}

	existPredictions = true

	for index, value := range records {
		param.Predictions[index] = make(map[string]int)
		param.Predictions[index]["Time_Id"] = value.Time_Id
		param.Predictions[index]["Number_1"] = value.Number_1
		param.Predictions[index]["Number_2"] = value.Number_2
		param.Predictions[index]["Number_3"] = value.Number_3
		param.Predictions[index]["Number_4"] = value.Number_4
		param.Predictions[index]["Number_5"] = value.Number_5
		param.Predictions[index]["Number_6"] = value.Number_6
		param.Predictions[index]["Number_7"] = value.Number_7
	}

	if beforeResultAnnouncement == true {
		param.Records = len(param.Predictions)
		return param, beforeResultAnnouncement, existPredictions
	}

	totalMatch := 0
	for i := 0; i < len(param.Predictions); i++ {
		param.Predictions[i]["Matches"] = 0
		if param.Predictions[i]["Number_1"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_1"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_1"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_1"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_1"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_1"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_1"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_2"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_3"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_4"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_5"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_6"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_1"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_2"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_3"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_4"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_5"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_6"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
		if param.Predictions[i]["Number_7"] == param.Result["Number_7"] {
			param.Predictions[i]["Matches"]++
			totalMatch++
		}
	}
	param.Records = len(param.Predictions)
	shift := math.Pow(10, 3)
	shift2 := math.Pow(10, 5)
	param.Average = math.Trunc((float64(totalMatch)/float64(param.Records))*shift) / shift
	param.Rate = math.Trunc((float64(totalMatch)/(float64(param.Records)*float64(7)))*shift2) / shift2 * float64(100)

	return param, beforeResultAnnouncement, existPredictions
}

func castSliceParamLoto7Detail(param Loto7UsersPredictionsDetailForCsv, beforeResultAnnouncement bool, existPredictions bool) ([][]interface{}, [][]interface{}) {
	sliceResult := make([][]interface{}, 0, 0)
	slicePredictions := make([][]interface{}, 0, 0)
	if beforeResultAnnouncement == false {
		sliceResult = append(sliceResult, []interface{}{param.Result["Number_1"], param.Result["Number_2"], param.Result["Number_3"], param.Result["Number_4"], param.Result["Number_5"], param.Result["Number_6"], param.Result["Number_7"]})
	}
	if existPredictions == true {
		for _, value := range param.Predictions {
			slicePredictions = append(slicePredictions, []interface{}{value["Time_Id"], value["Number_1"], value["Number_2"], value["Number_3"], value["Number_4"], value["Number_5"], value["Number_6"], value["Number_7"], value["Matches"]})
		}
	}
	return sliceResult, slicePredictions
}

func castParamLoto7Detail(time int, param Loto7UsersPredictionsDetailForCsv, sliceResult [][]interface{}, slicePredictions [][]interface{}, beforeResultAnnouncement bool, existPredictions bool) [][]string {
	castedParam := make([][]string, 0, 0)

	// Title
	castedParam = append(castedParam, []string{"Time : " + strconv.Itoa(time)})

	if beforeResultAnnouncement == false {
		// Title
		castedParam = append(castedParam, []string{"Result"})
		// Header
		castedParam = append(castedParam, []string{"Number1", "Number2", "Number3", "Number4", "Number5", "Number6", "Number7"})

		// Result
		castedParamResult := make([]string, 7)
		for i := 0; i < len(sliceResult[0]); i++ {
			castedParamResult[i] = strconv.Itoa(sliceResult[0][i].(int))
		}
		castedParam = append(castedParam, castedParamResult)

		// Blank
		castedParam = append(castedParam, []string{})
	}

	if existPredictions == true {
		// Title
		castedParam = append(castedParam, []string{"Predictions"})

		// Header
		if beforeResultAnnouncement == true {
			castedParam = append(castedParam, []string{"Records"})
		} else {
			castedParam = append(castedParam, []string{"Records", "Average", "Rate"})
		}

		// Data
		if beforeResultAnnouncement == true {
			castedParam = append(castedParam, []string{strconv.Itoa(param.Records)})
		} else {
			castedParamAfterStatistics := make([]string, 3)
			castedParamAfterStatistics[0] = strconv.Itoa(param.Records)
			castedParamAfterStatistics[1] = strconv.FormatFloat(param.Average, 'f', -1, 64)
			castedParamAfterStatistics[2] = strconv.FormatFloat(param.Rate, 'f', -1, 64) + "%"
			castedParam = append(castedParam, castedParamAfterStatistics)
		}

		// Header
		if beforeResultAnnouncement == true {
			castedParam = append(castedParam, []string{"Number1", "Number2", "Number3", "Number4", "Number5", "Number6", "Number7"})
		} else {
			castedParam = append(castedParam, []string{"Number1", "Number2", "Number3", "Number4", "Number5", "Number6", "Number7", "Matches"})
		}

		// Records
		for i := 0; i < len(slicePredictions); i++ {
			if beforeResultAnnouncement == true {
				castedParamPredictions := make([]string, 7)
				castedParamPredictions[0] = strconv.Itoa(slicePredictions[i][1].(int))
				castedParamPredictions[1] = strconv.Itoa(slicePredictions[i][2].(int))
				castedParamPredictions[2] = strconv.Itoa(slicePredictions[i][3].(int))
				castedParamPredictions[3] = strconv.Itoa(slicePredictions[i][4].(int))
				castedParamPredictions[4] = strconv.Itoa(slicePredictions[i][5].(int))
				castedParamPredictions[5] = strconv.Itoa(slicePredictions[i][6].(int))
				castedParamPredictions[6] = strconv.Itoa(slicePredictions[i][7].(int))
				castedParam = append(castedParam, castedParamPredictions)
			} else {
				castedParamPredictions := make([]string, 8)
				castedParamPredictions[0] = strconv.Itoa(slicePredictions[i][1].(int))
				castedParamPredictions[1] = strconv.Itoa(slicePredictions[i][2].(int))
				castedParamPredictions[2] = strconv.Itoa(slicePredictions[i][3].(int))
				castedParamPredictions[3] = strconv.Itoa(slicePredictions[i][4].(int))
				castedParamPredictions[4] = strconv.Itoa(slicePredictions[i][5].(int))
				castedParamPredictions[5] = strconv.Itoa(slicePredictions[i][6].(int))
				castedParamPredictions[6] = strconv.Itoa(slicePredictions[i][7].(int))
				castedParamPredictions[7] = strconv.Itoa(slicePredictions[i][8].(int))
				castedParam = append(castedParam, castedParamPredictions)
			}
		}
	}
	return castedParam
}
