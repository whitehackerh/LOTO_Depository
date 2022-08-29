package Services

import (
	"sort"

	model "../Models"
)

type Loto7LatelyStatistics struct {
	LatestTime1  int
	LatestTime2  int
	LatestTime3  int
	LatestTime4  int
	LatestTime5  int
	LatestTime6  int
	LatestTime7  int
	LatestTime8  int
	LatestTime9  int
	LatestTime10 int
	Rank         int
	Number       int
	Count        int
	Time1        int
	Time2        int
	Time3        int
	Time4        int
	Time5        int
	Time6        int
	Time7        int
	Time8        int
	Time9        int
	Time10       int
}

func GetLoto7LatelyStatistics() []Loto7LatelyStatistics {
	data := model.GetLoto7LatelyStatistics()
	lately_statistics := make([]Loto7LatelyStatistics, 38)
	lately_statistics[0].LatestTime1 = data[0].Time
	lately_statistics[0].LatestTime2 = data[1].Time
	lately_statistics[0].LatestTime3 = data[2].Time
	lately_statistics[0].LatestTime4 = data[3].Time
	lately_statistics[0].LatestTime5 = data[4].Time
	lately_statistics[0].LatestTime6 = data[5].Time
	lately_statistics[0].LatestTime7 = data[6].Time
	lately_statistics[0].LatestTime8 = data[7].Time
	lately_statistics[0].LatestTime9 = data[8].Time
	lately_statistics[0].LatestTime10 = data[9].Time
	counter := 0
	for _, value := range data {
		counter++
		for i := 1; i <= 37; i++ {
			lately_statistics[i].Number = i
			if value.Number_1 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
			if value.Number_2 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
			if value.Number_3 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
			if value.Number_4 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
			if value.Number_5 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
			if value.Number_6 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
			if value.Number_7 == i {
				switch counter {
				case 1:
					lately_statistics[i].Time1 = 1
					break
				case 2:
					lately_statistics[i].Time2 = 1
					break
				case 3:
					lately_statistics[i].Time3 = 1
					break
				case 4:
					lately_statistics[i].Time4 = 1
					break
				case 5:
					lately_statistics[i].Time5 = 1
					break
				case 6:
					lately_statistics[i].Time6 = 1
					break
				case 7:
					lately_statistics[i].Time7 = 1
					break
				case 8:
					lately_statistics[i].Time8 = 1
					break
				case 9:
					lately_statistics[i].Time9 = 1
					break
				case 10:
					lately_statistics[i].Time10 = 1
					break
				}
				lately_statistics[i].Count++
			}
		}
	}

	sort.Slice(lately_statistics, func(i, j int) bool {
		if lately_statistics[i].Count > lately_statistics[j].Count {
			return true
		} else if lately_statistics[i].Count == lately_statistics[j].Count {
			if lately_statistics[i].LatestTime1 < lately_statistics[j].LatestTime1 {
				return true
			} else if lately_statistics[i].LatestTime1 == lately_statistics[j].LatestTime1 {
				if lately_statistics[i].Number < lately_statistics[j].Number {
					return true
				}
			}
		}
		return false
	})

	lately_statistics = addRankLoto7(lately_statistics)
	return lately_statistics
}

func addRankLoto7(lately_statistics []Loto7LatelyStatistics) []Loto7LatelyStatistics {
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
