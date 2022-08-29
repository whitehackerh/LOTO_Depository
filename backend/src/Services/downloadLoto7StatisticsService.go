package Services

import (
	"log"
	"os"

	model "../Models"
	"github.com/gocarina/gocsv"
)

func DownloadLoto7Statistics() {
	data := model.GetLoto7StatisticsCsv()

	file, err := os.OpenFile("Loto7Statistics.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	gocsv.MarshalFile(data, file)
}
