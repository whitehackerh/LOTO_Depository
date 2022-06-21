package Model

import db "../DB"

type Loto7Statistics struct {
	Rank   int
	Number int
	Count  int
	Rate   float64
	Time   int
}

type Loto7StatisticsCsv struct {
	Rank   int
	Number int
	Count  int
	Rate   float64
}

func init() {
	Db = db.ConnectDb()
}

func GetLoto7Statistics() []*Loto7Statistics {
	statistics := Loto7Statistics{}
	data := []*Loto7Statistics{}
	rows, err := Db.Query("SELECT RANK() OVER(ORDER BY count DESC) AS Rank, Number, Count, Rate, Time FROM loto7_statistics ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Rank, &statistics.Number, &statistics.Count, &statistics.Rate, &statistics.Time)
		data = append(data, &Loto7Statistics{Rank: statistics.Rank, Number: statistics.Number, Count: statistics.Count, Rate: statistics.Rate, Time: statistics.Time})
	}
	return data
}

func GetLoto7StatisticsCsv() []*Loto7StatisticsCsv {
	statistics := Loto7StatisticsCsv{}
	data := []*Loto7StatisticsCsv{}
	rows, err := Db.Query("SELECT RANK() OVER(ORDER BY count DESC) AS Rank, Number, Count, Rate FROM loto7_statistics ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Rank, &statistics.Number, &statistics.Count, &statistics.Rate)
		data = append(data, &Loto7StatisticsCsv{Rank: statistics.Rank, Number: statistics.Number, Count: statistics.Count, Rate: statistics.Rate})
	}
	return data
}
