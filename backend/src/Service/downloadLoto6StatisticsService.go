package Service

import (
	"log"
	"os"

	model "../Model"
	"github.com/gocarina/gocsv"
)

func DownloadLoto6Statistics() {
	data := model.GetLoto6StatisticsCsv()

	file, err := os.OpenFile("Loto6Statistics.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	gocsv.MarshalFile(data, file)
}
