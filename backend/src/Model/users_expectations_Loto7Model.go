package Model

import (
	db "../DB"
)

func init() {
	Db = db.ConnectDb()
}

func SetLoto7Expectations(input_data map[int]map[string]int, user_id int) bool {
	query := `INSERT INTO users_expectations_loto7(user_id, time, time_id, number_1, number_2, number_3, number_4, number_5, number_6, number_7) VALUES($1, $2, (SELECT CASE count(*) WHEN 0 THEN 1 ELSE MAX(time_id) + 1 END FROM users_expectations_loto7 WHERE user_id = $3 AND time = $4), $5, $6, $7, $8, $9, $10, $11)`
	for i := 0; i < len(input_data); i++ {
		_, err := Db.Exec(query, user_id, input_data[i]["time"], user_id, input_data[i]["time"], input_data[i]["input_number_1"], input_data[i]["input_number_2"], input_data[i]["input_number_3"], input_data[i]["input_number_4"], input_data[i]["input_number_5"], input_data[i]["input_number_6"], input_data[i]["input_number_7"])
		if err != nil {
			return false
		}
	}
	return true
}
