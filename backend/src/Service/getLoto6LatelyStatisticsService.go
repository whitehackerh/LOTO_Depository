package Service

import (
	"sort"

	model "../Model"
)

type Loto6LatelyStatistics struct {
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

func GetLoto6LatelyStatistics() []Loto6LatelyStatistics {
	data := model.GetLoto6LatelyStatistics()
	lately_statistics := make([]Loto6LatelyStatistics, 44)
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
	count := 0
	for _, value := range data {
		count++
		for i := 1; i <= 43; i++ {
			lately_statistics[i].Number = i
			if value.Number_1 == i {
				switch count {
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
				switch count {
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
				switch count {
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
				switch count {
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
				switch count {
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
				switch count {
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

	lately_statistics = addRank(lately_statistics)
	return lately_statistics
}

func addRank(lately_statistics []Loto6LatelyStatistics) []Loto6LatelyStatistics {
	for i := 0; i <= 43; i++ {
		if i <= 42 {
			lately_statistics[i].Rank = i + 1
		}
	}
	return lately_statistics
}
