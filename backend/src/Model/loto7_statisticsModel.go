package Model

import db "../DB"

type Loto7Statistics struct {
	Number int
	Time   int
	Count  int
	Rate   float64
}

func init() {
	Db = db.ConnectDb()
}

func GetLoto7Statistics() []*Loto7Statistics {
	statistics := Loto7Statistics{}
	data := []*Loto7Statistics{}
	rows, err := Db.Query("SELECT * FROM Loto7_statistics ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&statistics.Number, &statistics.Time, &statistics.Count, &statistics.Rate)
		data = append(data, &Loto7Statistics{Number: statistics.Number, Time: statistics.Time, Count: statistics.Count, Rate: statistics.Rate})
	}
	return data
}
