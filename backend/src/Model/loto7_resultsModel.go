package Model

import (
	"fmt"

	db "../DB"

	_ "github.com/lib/pq"
)

type Loto7Results struct {
	Time     int
	Number_1 int
	Number_2 int
	Number_3 int
	Number_4 int
	Number_5 int
	Number_6 int
	Number_7 int
}

func init() {
	Db = db.ConnectDb()
}

func GetLoto7Results() []*Loto7Results {
	result := Loto7Results{}
	data := []*Loto7Results{}

	rows, err := Db.Query("select * from loto7_results order by time desc")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&result.Time, &result.Number_1, &result.Number_2, &result.Number_3,
			&result.Number_4, &result.Number_5, &result.Number_6, &result.Number_7)
		data = append(data, &Loto7Results{Time: result.Time, Number_1: result.Number_1, Number_2: result.Number_2, Number_3: result.Number_3,
			Number_4: result.Number_4, Number_5: result.Number_5, Number_6: result.Number_6, Number_7: result.Number_7})
	}
	return data
}

func SetLoto7Results(input_data map[string]int) bool {
	query := `INSERT INTO loto7_results(time, number_1, number_2, number_3, number_4, number_5, number_6, number_7) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`
	res, err := Db.Exec(query, input_data["time"], input_data["input_number_1"], input_data["input_number_2"], input_data["input_number_3"], input_data["input_number_4"], input_data["input_number_5"], input_data["input_number_6"], input_data["input_number_7"])
	fmt.Println(res)
	if err != nil {
		return false
	}
	return true
}
