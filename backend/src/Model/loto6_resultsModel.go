package Model

import (
	"fmt"

	db "../DB"
	_ "github.com/lib/pq"
)

type Loto6Results struct {
	Time             int
	Winning_number_1 int
	Winning_number_2 int
	Winning_number_3 int
	Winning_number_4 int
	Winning_number_5 int
	Winning_number_6 int
}

func init() {
	Db = db.ConnectDb()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func GetLoto6Results() []*Loto6Results {
	result := Loto6Results{}
	data := []*Loto6Results{}
	rows, err := Db.Query("select * from loto6_results order by time desc")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result.Time, &result.Winning_number_1, &result.Winning_number_2, &result.Winning_number_3,
			&result.Winning_number_4, &result.Winning_number_5, &result.Winning_number_6)
		data = append(data, &Loto6Results{Time: result.Time, Winning_number_1: result.Winning_number_1, Winning_number_2: result.Winning_number_2, Winning_number_3: result.Winning_number_3,
			Winning_number_4: result.Winning_number_4, Winning_number_5: result.Winning_number_5, Winning_number_6: result.Winning_number_6})
	}
	return data
}

func SetLoto6Results(data map[string]int) bool {
	query := `INSERT INTO loto6_results(time, winning_number_1, winning_number_2, winning_number_3, winning_number_4, winning_number_5, winning_number_6) VALUES($1, $2, $3, $4, $5, $6, $7);`
	res, err := Db.Exec(query, data["time"], data["data_1"], data["data_2"], data["data_3"], data["data_4"], data["data_5"], data["data_6"])
	fmt.Println(res)
	if err != nil {
		return false
	}
	return true
}
