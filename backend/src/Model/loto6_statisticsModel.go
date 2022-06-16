package Model

import db "../DB"

type Loto6Statistics struct {
	Number int
	Time   int
	Count  int
	Rate   float64
}

func init() {
	Db = db.ConnectDb()
}

func GetLoto6Statistics() []*Loto6Statistics {
	statistics := Loto6Statistics{}
	data := []*Loto6Statistics{}
	rows, err := Db.Query("SELECT * FROM loto6_statistics ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Number, &statistics.Time, &statistics.Count, &statistics.Rate)
		data = append(data, &Loto6Statistics{Number: statistics.Number, Time: statistics.Time, Count: statistics.Count, Rate: statistics.Rate})
	}
	return data
}
