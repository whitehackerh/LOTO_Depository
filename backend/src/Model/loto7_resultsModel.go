package Model

import (
	"strconv"

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

type Loto7LatelyResults struct {
	Time     int
	Number_1 int
	Number_2 int
	Number_3 int
	Number_4 int
	Number_5 int
	Number_6 int
	Number_7 int
}

type Loto7ResultDetail struct {
	Time     int
	Number_1 int
	Number_2 int
	Number_3 int
	Number_4 int
	Number_5 int
	Number_6 int
	Number_7 int
}

type Loto7NewestTime struct {
	Time int
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

func SetLoto7Results(input_data map[string]int, input_row [7]string) bool {
	query := `INSERT INTO loto7_results(time, number_1, number_2, number_3, number_4, number_5, number_6, number_7) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := Db.Exec(query, input_data["time"], input_data["input_number_1"], input_data["input_number_2"], input_data["input_number_3"], input_data["input_number_4"], input_data["input_number_5"], input_data["input_number_6"], input_data["input_number_7"])
	if err != nil {
		return false
	}

	query2 := `UPDATE loto7_statistics SET time = (SELECT MAX(loto7_results.time) FROM loto7_results)`
	_, err2 := Db.Exec(query2)
	if err2 != nil {
		return false
	}

	for i := 0; i < len(input_row); i++ {
		query3 := `UPDATE loto7_statistics SET count = (SELECT count + 1 FROM loto7_statistics WHERE number = ` + input_row[i] + ` ) WHERE number = ` + input_row[i]
		_, err3 := Db.Exec(query3)
		if err3 != nil {
			return false
		}
	}

	for i := 1; i <= 37; i++ {
		castedCounter := strconv.Itoa(i)
		query4 := `UPDATE loto7_statistics SET rate = (SELECT CAST(count AS float) FROM loto7_statistics WHERE number = ` + castedCounter + ` ) / (SELECT time FROM loto7_statistics WHERE number = ` + castedCounter + ` ) WHERE number = ` + castedCounter
		_, err4 := Db.Exec(query4)
		if err4 != nil {
			return false
		}
	}

	return true
}

func GetLoto7LatelyStatistics() []*Loto7LatelyResults {
	result := Loto7LatelyResults{}
	data := []*Loto7LatelyResults{}
	rows, err := Db.Query("select * from loto7_results order by time desc LIMIT 10")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result.Time, &result.Number_1, &result.Number_2, &result.Number_3,
			&result.Number_4, &result.Number_5, &result.Number_6, &result.Number_7)
		data = append(data, &Loto7LatelyResults{Time: result.Time, Number_1: result.Number_1, Number_2: result.Number_2, Number_3: result.Number_3,
			Number_4: result.Number_4, Number_5: result.Number_5, Number_6: result.Number_6, Number_7: result.Number_7})
	}
	return data
}

func GetLoto7ResultDetail(time int) []*Loto7ResultDetail {
	result := Loto7ResultDetail{}
	data := []*Loto7ResultDetail{}
	query := `select * from loto7_results where time = $1`
	rows, err := Db.Query(query, time)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result.Time, &result.Number_1, &result.Number_2, &result.Number_3,
			&result.Number_4, &result.Number_5, &result.Number_6, &result.Number_7)
		data = append(data, &Loto7ResultDetail{Time: result.Time, Number_1: result.Number_1, Number_2: result.Number_2, Number_3: result.Number_3,
			Number_4: result.Number_4, Number_5: result.Number_5, Number_6: result.Number_6, Number_7: result.Number_7})
	}
	return data
}

func UpdateLoto7Results(updateLoto7ResultsQuery string, input_data map[string]int, updateStatisticsQuery map[int]string, params map[int]int, time int) bool {
	_, err := Db.Exec(updateLoto7ResultsQuery, input_data["number_1"], input_data["number_2"], input_data["number_3"], input_data["number_4"], input_data["number_5"], input_data["number_6"], input_data["number_7"], time)
	if err != nil {
		return false
	}

	for i := 0; i < len(updateStatisticsQuery); i++ {
		_, err := Db.Exec(updateStatisticsQuery[i], params[i], params[i], params[i], params[i])
		if err != nil {
			return false
		}
	}
	return true
}

func GetNewestLoto7Result() []*Loto7NewestTime {
	result := Loto7NewestTime{}
	data := []*Loto7NewestTime{}

	rows, err := Db.Query("SELECT time FROM loto7_results ORDER BY time DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result.Time)
		data = append(data, &Loto7NewestTime{Time: result.Time})
	}
	return data
}
