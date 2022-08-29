package Services

import (
	"sort"
	"strconv"

	model "../Models"
)

type Loto7LatelyStatisticsForCsv struct {
	Rank   int
	Number int
	Count  int
	Time1  string
	Time2  string
	Time3  string
	Time4  string
	Time5  string
	Time6  string
	Time7  string
	Time8  string
	Time9  string
	Time10 string
}

func DownloadLoto7LatelyStatistics() {
	data := model.GetLoto7LatelyStatistics()
	lately_statistics := make([]Loto7LatelyStatisticsForCsv, 37)
	timeList := []string{"Rank", "Number", "Count", strconv.Itoa(data[0].Time), strconv.Itoa(data[1].Time), strconv.Itoa(data[2].Time), strconv.Itoa(data[3].Time), strconv.Itoa(data[4].Time),
		strconv.Itoa(data[5].Time), strconv.Itoa(data[6].Time), strconv.Itoa(data[7].Time), strconv.Itoa(data[8].Time), strconv.Itoa(data[9].Time)}
	counter := 0
	for _, value := range data {
		counter++
		for i := 1; i <= 37; i++ {
			lately_statistics[i-1].Number = i
			if value.Number_1 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
			if value.Number_2 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
			if value.Number_3 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
			if value.Number_4 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
			if value.Number_5 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
			if value.Number_6 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
			if value.Number_7 == i {
				switch counter {
				case 1:
					lately_statistics[i-1].Time1 = "○"
					break
				case 2:
					lately_statistics[i-1].Time2 = "○"
					break
				case 3:
					lately_statistics[i-1].Time3 = "○"
					break
				case 4:
					lately_statistics[i-1].Time4 = "○"
					break
				case 5:
					lately_statistics[i-1].Time5 = "○"
					break
				case 6:
					lately_statistics[i-1].Time6 = "○"
					break
				case 7:
					lately_statistics[i-1].Time7 = "○"
					break
				case 8:
					lately_statistics[i-1].Time8 = "○"
					break
				case 9:
					lately_statistics[i-1].Time9 = "○"
					break
				case 10:
					lately_statistics[i-1].Time10 = "○"
					break
				}
				lately_statistics[i-1].Count++
			}
		}
	}

	lately_statistics = convertStringLoto7(lately_statistics)
	sort.Slice(lately_statistics, func(i, j int) bool {
		if lately_statistics[i].Count > lately_statistics[j].Count {
			return true
		} else if lately_statistics[i].Count == lately_statistics[j].Count {
			if lately_statistics[i].Number < lately_statistics[j].Number {
				return true
			}
		}
		return false
	})
	lately_statistics = addRankForLoto7Csv(lately_statistics)
	line := make([][]string, 37)
	for i := 0; i <= 36; i++ {
		line[i] = make([]string, 13)
		for j := 0; j <= 12; j++ {
			switch j {
			case 0:
				line[i][0] = strconv.Itoa(lately_statistics[i].Rank)
				break
			case 1:
				line[i][1] = strconv.Itoa(lately_statistics[i].Number)
				break
			case 2:
				line[i][2] = strconv.Itoa(lately_statistics[i].Count)
				break
			case 3:
				line[i][3] = lately_statistics[i].Time1
				break
			case 4:
				line[i][4] = lately_statistics[i].Time2
				break
			case 5:
				line[i][5] = lately_statistics[i].Time3
				break
			case 6:
				line[i][6] = lately_statistics[i].Time4
				break
			case 7:
				line[i][7] = lately_statistics[i].Time5
				break
			case 8:
				line[i][8] = lately_statistics[i].Time6
				break
			case 9:
				line[i][9] = lately_statistics[i].Time7
				break
			case 10:
				line[i][10] = lately_statistics[i].Time8
				break
			case 11:
				line[i][11] = lately_statistics[i].Time9
				break
			case 12:
				line[i][12] = lately_statistics[i].Time10
				break
			}
		}
	}
	Write_csv_col_header("Loto7LatelyStatistics.csv", timeList, line)
}

func convertStringLoto7(lately_statistics []Loto7LatelyStatisticsForCsv) []Loto7LatelyStatisticsForCsv {
	for i := 0; i < 37; i++ {
		if lately_statistics[i].Time1 == "" {
			lately_statistics[i].Time1 = "---"
		}
		if lately_statistics[i].Time2 == "" {
			lately_statistics[i].Time2 = "---"
		}
		if lately_statistics[i].Time3 == "" {
			lately_statistics[i].Time3 = "---"
		}
		if lately_statistics[i].Time4 == "" {
			lately_statistics[i].Time4 = "---"
		}
		if lately_statistics[i].Time5 == "" {
			lately_statistics[i].Time5 = "---"
		}
		if lately_statistics[i].Time6 == "" {
			lately_statistics[i].Time6 = "---"
		}
		if lately_statistics[i].Time7 == "" {
			lately_statistics[i].Time7 = "---"
		}
		if lately_statistics[i].Time8 == "" {
			lately_statistics[i].Time8 = "---"
		}
		if lately_statistics[i].Time9 == "" {
			lately_statistics[i].Time9 = "---"
		}
		if lately_statistics[i].Time10 == "" {
			lately_statistics[i].Time10 = "---"
		}
	}
	return lately_statistics
}

func addRankForLoto7Csv(lately_statistics []Loto7LatelyStatisticsForCsv) []Loto7LatelyStatisticsForCsv {
	for i := 0; i <= 36; i++ {
		lately_statistics[i].Rank = 1
		for j := 0; j <= 36; j++ {
			if lately_statistics[j].Count > lately_statistics[i].Count {
				lately_statistics[i].Rank++
			}
		}
	}
	return lately_statistics
}
