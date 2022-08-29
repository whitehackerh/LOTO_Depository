package Models

import db "../DB"

type Loto6Statistics struct {
	Rank   int
	Number int
	Count  int
	Rate   float64
	Time   int
}

type Loto6StatisticsCsv struct {
	Rank   int
	Number int
	Count  int
	Rate   float64
}

func init() {
	Db = db.ConnectDb()
}

func GetLoto6Statistics() []*Loto6Statistics {
	statistics := Loto6Statistics{}
	data := []*Loto6Statistics{}
	rows, err := Db.Query("SELECT RANK() OVER(ORDER BY count DESC) AS Rank, Number, Count, Rate, Time FROM loto6_statistics ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Rank, &statistics.Number, &statistics.Count, &statistics.Rate, &statistics.Time)
		data = append(data, &Loto6Statistics{Rank: statistics.Rank, Number: statistics.Number, Count: statistics.Count, Rate: statistics.Rate, Time: statistics.Time})
	}
	return data
}

func GetLoto6StatisticsCsv() []*Loto6StatisticsCsv {
	statistics := Loto6StatisticsCsv{}
	data := []*Loto6StatisticsCsv{}
	rows, err := Db.Query("SELECT RANK() OVER(ORDER BY count DESC) AS Rank, Number, Count, Rate FROM loto6_statistics ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Rank, &statistics.Number, &statistics.Count, &statistics.Rate)
		data = append(data, &Loto6StatisticsCsv{Rank: statistics.Rank, Number: statistics.Number, Count: statistics.Count, Rate: statistics.Rate})
	}
	return data
}

func GetLoto6StatisticsDetail(numbers map[string]int) []*Loto6Statistics {
	statistics := Loto6Statistics{}
	data := []*Loto6Statistics{}
	query := `SELECT * FROM loto6_statistics WHERE number = $1 or number = $2 or number = $3 or number = $4 or number = $5 or number = $6`
	rows, err := Db.Query(query, numbers["number_1"], numbers["number_2"], numbers["number_3"], numbers["number_4"], numbers["number_5"], numbers["number_6"])
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Rank, &statistics.Number, &statistics.Count, &statistics.Rate, &statistics.Time)
		data = append(data, &Loto6Statistics{Rank: statistics.Rank, Number: statistics.Number, Count: statistics.Count, Rate: statistics.Rate, Time: statistics.Time})
	}
	return data
}
