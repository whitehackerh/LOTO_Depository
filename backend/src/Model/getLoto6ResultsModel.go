package Model

import (
	"database/sql"

	// "net/http"

	// "github.com/labstack/echo"

	_ "github.com/lib/pq"
	// "github.com/pkg/errors"
)

// type Results struct {
// 	Time             int `json:"time"`
// 	Winning_number_1 int `json:"winning_number_1"`
// 	Winning_number_2 int `json:"winning_number_2"`
// 	Winning_number_3 int `json:"winning_number_3"`
// 	Winning_number_4 int `json:"winning_number_4"`
// 	Winning_number_5 int `json:"winning_number_5"`
// 	Winning_number_6 int `json:"winning_number_6"`
// }

type Loto6Results struct {
	Time             int
	Winning_number_1 int
	Winning_number_2 int
	Winning_number_3 int
	Winning_number_4 int
	Winning_number_5 int
	Winning_number_6 int
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=loto_depository password=ash0001 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// func GetLoto6Results() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		result := Results{}
// 		results := []*Results{}

// 		rows, err := Db.Query("select * from loto6_results order by time")
// 		if err != nil {
// 			return errors.Wrapf(err, "cannot connect SQL")
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			if err := rows.Scan(&result.Time, &result.Winning_number_1, &result.Winning_number_2, &result.Winning_number_3,
// 				&result.Winning_number_4, &result.Winning_number_5, &result.Winning_number_6); err != nil {
// 				return errors.Wrapf(err, "cannot connect SQL")
// 			}
// 			results = append(results, &Results{Time: result.Time, Winning_number_1: result.Winning_number_1, Winning_number_2: result.Winning_number_2, Winning_number_3: result.Winning_number_3,
// 				Winning_number_4: result.Winning_number_4, Winning_number_5: result.Winning_number_5, Winning_number_6: result.Winning_number_6})
// 		}
// 		return c.JSON(http.StatusOK, results)
// 	}
// }

// type Results []Results

func GetLoto6Results() []*Loto6Results {
	result := Loto6Results{}
	data := []*Loto6Results{}
	//results := Results{}
	//record := make(map[string]int)
	// record := make(map[string]int)
	// var data []map[string]int
	//data := []record

	rows, err := Db.Query("select * from loto6_results order by time")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// for rows.Next() {

	// 	// rows.Scan(&result.Time, &result.Winning_number_1, &result.Winning_number_2, &result.Winning_number_3,
	// 	// 	&result.Winning_number_4, &result.Winning_number_5, &result.Winning_number_6)
	// 	// //fmt.Println(data)
	// 	// //results = append(results, result)
	// 	// record["time"] = result.Time
	// 	// fmt.Println(data)
	// 	// record["winning_number_1"] = result.Winning_number_1
	// 	// record["winning_number_2"] = result.Winning_number_2
	// 	// record["winning_number_3"] = result.Winning_number_3
	// 	// record["winning_number_4"] = result.Winning_number_4
	// 	// record["winning_number_5"] = result.Winning_number_5
	// 	// record["winning_number_6"] = result.Winning_number_6
	// 	// fmt.Println(record)
	// 	// fmt.Println(data)
	// 	// data = append(data, record)
	// 	// fmt.Println(data)
	// 	// data = append(data, &record{"time": result.Time, "winning_number_1": result.Winning_number_1, "winning_number_2": result.Winning_number_2, "winning_number_3": result.Winning_number_3,
	// 	// 	"winning_number_4": result.Winning_number_4, "winning_number_5": result.Winning_number_5, "winning_number_6": result.Winning_number_6})

	// }
	for rows.Next() {
		rows.Scan(&result.Time, &result.Winning_number_1, &result.Winning_number_2, &result.Winning_number_3,
			&result.Winning_number_4, &result.Winning_number_5, &result.Winning_number_6)
		data = append(data, &Loto6Results{Time: result.Time, Winning_number_1: result.Winning_number_1, Winning_number_2: result.Winning_number_2, Winning_number_3: result.Winning_number_3,
			Winning_number_4: result.Winning_number_4, Winning_number_5: result.Winning_number_5, Winning_number_6: result.Winning_number_6})
	}
	return data
}
