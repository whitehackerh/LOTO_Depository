package Model

import (
	"fmt"

	db "../DB"

	_ "github.com/lib/pq"
)

type Loto7Results struct {
	Time             int
	Winning_number_1 int
	Winning_number_2 int
	Winning_number_3 int
	Winning_number_4 int
	Winning_number_5 int
	Winning_number_6 int
	Winning_number_7 int
}

func init() {
	Db = db.ConnectDb()
}

func GetLoto7Results() []*Loto7Results {
	result := Loto7Results{}
	data := []*Loto7Results{}

	rows, err := Db.Query("select * from loto7_results order by time")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&result.Time, &result.Winning_number_1, &result.Winning_number_2, &result.Winning_number_3,
			&result.Winning_number_4, &result.Winning_number_5, &result.Winning_number_6, &result.Winning_number_7)
		data = append(data, &Loto7Results{Time: result.Time, Winning_number_1: result.Winning_number_1, Winning_number_2: result.Winning_number_2, Winning_number_3: result.Winning_number_3,
			Winning_number_4: result.Winning_number_4, Winning_number_5: result.Winning_number_5, Winning_number_6: result.Winning_number_6, Winning_number_7: result.Winning_number_7})
	}
	return data
}

func SetLoto7Results(data map[string]int) bool {
	query := `INSERT INTO loto7_results(time, winning_number_1, winning_number_2, winning_number_3, winning_number_4, winning_number_5, winning_number_6, winning_number_7) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`
	res, err := Db.Exec(query, data["time"], data["data_1"], data["data_2"], data["data_3"], data["data_4"], data["data_5"], data["data_6"], data["data_7"])
	fmt.Println(res)
	if err != nil {
		return false
	}
	return true
}
