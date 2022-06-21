package Service

import (
	"log"
	"os"

	model "../Model"
	"github.com/gocarina/gocsv"
)

func DownloadLoto7Results() {
	data := model.GetLoto7Results()

	file, err := os.OpenFile("Loto7Results.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln("Error", err)
	}
	defer file.Close()

	gocsv.MarshalFile(data, file)
}
