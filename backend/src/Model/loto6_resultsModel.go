package Model

import (
	"fmt"

	db "../DB"
	_ "github.com/lib/pq"
)

type Loto6Results struct {
	Time     int
	Number_1 int
	Number_2 int
	Number_3 int
	Number_4 int
	Number_5 int
	Number_6 int
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
		rows.Scan(&result.Time, &result.Number_1, &result.Number_2, &result.Number_3,
			&result.Number_4, &result.Number_5, &result.Number_6)
		data = append(data, &Loto6Results{Time: result.Time, Number_1: result.Number_1, Number_2: result.Number_2, Number_3: result.Number_3,
			Number_4: result.Number_4, Number_5: result.Number_5, Number_6: result.Number_6})
	}
	fmt.Println(data)
	return data
}

func SetLoto6Results(input_data map[string]int) bool {
	query := `INSERT INTO loto6_results(time, number_1, number_2, number_3, number_4, number_5, number_6) VALUES($1, $2, $3, $4, $5, $6, $7);`
	res, err := Db.Exec(query, input_data["time"], input_data["input_number_1"], input_data["input_number_2"], input_data["input_number_3"], input_data["input_number_4"], input_data["input_number_5"], input_data["input_number_6"])
	fmt.Println(res)
	if err != nil {
		return false
	}
	return true
}
