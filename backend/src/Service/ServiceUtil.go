package Service

import (
	"encoding/csv"
	"os"
)

func Write_csv_col_header(filepath string, header []string, data [][]string) {
	csvfile, _ := os.Create(filepath)
	writer := csv.NewWriter(csvfile)
	writer.Write(header)
	for _, line := range data {
		writer.Write(line)
	}
	writer.Flush()
}
