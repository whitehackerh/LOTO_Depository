package Model

import (
	"fmt"

	db "../DB"
	_ "github.com/lib/pq"
)

type Loto6UsersPredictions struct {
	User_Id  int
	Time     int
	Time_Id  int
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

func SetLoto6Predictions(input_data map[int]map[string]int, user_id int) bool {
	query := `INSERT INTO users_expectations_loto6(user_id, time, time_id, number_1, number_2, number_3, number_4, number_5, number_6) VALUES($1, $2, (SELECT CASE count(*) WHEN 0 THEN 1 ELSE MAX(time_id) + 1 END FROM users_expectations_loto6 WHERE user_id = $3 AND time = $4), $5, $6, $7, $8, $9, $10)`
	for i := 0; i < len(input_data); i++ {
		_, err := Db.Exec(query, user_id, input_data[i]["time"], user_id, input_data[i]["time"], input_data[i]["input_number_1"], input_data[i]["input_number_2"], input_data[i]["input_number_3"], input_data[i]["input_number_4"], input_data[i]["input_number_5"], input_data[i]["input_number_6"])
		if err != nil {
			return false
		}
	}
	return true
}

func GetLoto6UsersPredictions(user_id int) []*Loto6UsersPredictions {
	prediction := Loto6UsersPredictions{}
	predictions := []*Loto6UsersPredictions{}
	var count int
	error := Db.QueryRow(`SELECT COUNT(*) FROM users_expectations_loto6 WHERE user_id = $1`, user_id).Scan(&count)
	fmt.Println(error)
	if count == 0 {
		prediction1 := Loto6UsersPredictions{Time: 0}
		predictions = append(predictions, &prediction1)
		return predictions
	}
	query := `SELECT * FROM users_expectations_loto6 WHERE user_id = $1 ORDER BY time DESC, time_id ASC`
	rows, err := Db.Query(query, user_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&prediction.User_Id, &prediction.Time, &prediction.Time_Id, &prediction.Number_1, &prediction.Number_2,
			&prediction.Number_3, &prediction.Number_4, &prediction.Number_5, &prediction.Number_6)
		predictions = append(predictions, &Loto6UsersPredictions{User_Id: prediction.User_Id, Time: prediction.Time, Time_Id: prediction.Time_Id, Number_1: prediction.Number_1, Number_2: prediction.Number_2,
			Number_3: prediction.Number_3, Number_4: prediction.Number_4, Number_5: prediction.Number_5, Number_6: prediction.Number_6})
	}
	return predictions
}
