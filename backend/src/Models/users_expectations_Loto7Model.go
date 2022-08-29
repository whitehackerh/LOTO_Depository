package Models

import (
	"fmt"

	db "../DB"
)

type Loto7UsersPredictions struct {
	User_Id  int
	Time     int
	Time_Id  int
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

func SetLoto7Predictions(input_data map[int]map[string]int, user_id int) bool {
	query := `INSERT INTO users_expectations_loto7(user_id, time, time_id, number_1, number_2, number_3, number_4, number_5, number_6, number_7) VALUES($1, $2, (SELECT CASE count(*) WHEN 0 THEN 1 ELSE MAX(time_id) + 1 END FROM users_expectations_loto7 WHERE user_id = $3 AND time = $4), $5, $6, $7, $8, $9, $10, $11)`
	for i := 0; i < len(input_data); i++ {
		_, err := Db.Exec(query, user_id, input_data[i]["time"], user_id, input_data[i]["time"], input_data[i]["input_number_1"], input_data[i]["input_number_2"], input_data[i]["input_number_3"], input_data[i]["input_number_4"], input_data[i]["input_number_5"], input_data[i]["input_number_6"], input_data[i]["input_number_7"])
		if err != nil {
			return false
		}
	}
	return true
}

func GetLoto7UsersPredictions(user_id int) []*Loto7UsersPredictions {
	prediction := Loto7UsersPredictions{}
	predictions := []*Loto7UsersPredictions{}
	var count int
	error := Db.QueryRow(`SELECT COUNT(*) FROM users_expectations_loto7 WHERE user_id = $1`, user_id).Scan(&count)
	fmt.Println(error)
	if count == 0 {
		prediction1 := Loto7UsersPredictions{Time: 0}
		predictions = append(predictions, &prediction1)
		return predictions
	}
	query := `SELECT * FROM users_expectations_loto7 WHERE user_id = $1 ORDER BY time DESC, time_id ASC`
	rows, err := Db.Query(query, user_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&prediction.User_Id, &prediction.Time, &prediction.Time_Id, &prediction.Number_1, &prediction.Number_2,
			&prediction.Number_3, &prediction.Number_4, &prediction.Number_5, &prediction.Number_6, &prediction.Number_7)
		predictions = append(predictions, &Loto7UsersPredictions{User_Id: prediction.User_Id, Time: prediction.Time, Time_Id: prediction.Time_Id, Number_1: prediction.Number_1, Number_2: prediction.Number_2,
			Number_3: prediction.Number_3, Number_4: prediction.Number_4, Number_5: prediction.Number_5, Number_6: prediction.Number_6, Number_7: prediction.Number_7})
	}
	return predictions
}

func GetLoto7UsersPredictionsDetail(user_id, time int) []*Loto7UsersPredictions {
	prediction := Loto7UsersPredictions{}
	predictions := []*Loto7UsersPredictions{}
	var count int
	error := Db.QueryRow(`SELECT COUNT(*) FROM users_expectations_loto7 WHERE user_id = $1 AND time = $2`, user_id, time).Scan(&count)
	fmt.Println(error)
	if count == 0 {
		prediction1 := Loto7UsersPredictions{Time: 0}
		predictions = append(predictions, &prediction1)
		return predictions
	}
	query := `SELECT * FROM users_expectations_loto7 WHERE user_id = $1 AND time = $2 ORDER BY time DESC, time_id ASC`
	rows, err := Db.Query(query, user_id, time)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&prediction.User_Id, &prediction.Time, &prediction.Time_Id, &prediction.Number_1, &prediction.Number_2,
			&prediction.Number_3, &prediction.Number_4, &prediction.Number_5, &prediction.Number_6, &prediction.Number_7)
		predictions = append(predictions, &Loto7UsersPredictions{User_Id: prediction.User_Id, Time: prediction.Time, Time_Id: prediction.Time_Id, Number_1: prediction.Number_1, Number_2: prediction.Number_2,
			Number_3: prediction.Number_3, Number_4: prediction.Number_4, Number_5: prediction.Number_5, Number_6: prediction.Number_6, Number_7: prediction.Number_7})
	}
	return predictions
}

func EditLoto7UsersPredictionsDetail(input_data map[int]map[string]int, user_id int) bool {
	query := `UPDATE users_expectations_loto7 SET number_1 = $1, number_2 = $2, number_3 = $3, number_4 = $4, number_5 = $5, number_6 = $6, number_7 = $7 WHERE user_id = $8 AND time = $9 AND time_id = $10`
	for i := 0; i < len(input_data); i++ {
		_, err := Db.Exec(query, input_data[i]["input_number_1"], input_data[i]["input_number_2"], input_data[i]["input_number_3"], input_data[i]["input_number_4"], input_data[i]["input_number_5"], input_data[i]["input_number_6"], input_data[i]["input_number_7"], user_id, input_data[i]["time"], input_data[i]["time_id"])
		if err != nil {
			return false
		}
	}
	return true
}

func DeleteLoto7UsersPredictionsDetail(delete_data map[int]map[string]int, user_id int) bool {
	query := `DELETE FROM users_expectations_loto7 WHERE user_id = $1 AND time = $2 AND time_id = $3`
	for i := 0; i < len(delete_data); i++ {
		_, err := Db.Exec(query, user_id, delete_data[i]["time"], delete_data[i]["time_id"])
		if err != nil {
			return false
		}
	}
	return true
}
