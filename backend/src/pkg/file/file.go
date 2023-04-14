package file

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func OpenCSV(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error: sasad")
		fmt.Println("Error: ", err)
	}
	return file
}

type Record [][]string

func ReadCSV(file *os.File) Record {
	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
