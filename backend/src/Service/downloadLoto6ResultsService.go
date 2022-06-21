package Service

import (
	"log"
	"os"

	model "../Model"
	"github.com/gocarina/gocsv"
)

func DownloadLoto6Results() {
	data := model.GetLoto6Results()

	file, err := os.OpenFile("Loto6Results.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	// Writing
	gocsv.MarshalFile(data, file)
}
