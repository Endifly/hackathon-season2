package main

import (
	"fmt"
	"os"
	"time"

	"github.com/markkj/hackathon-season2/internal/csv"
	"github.com/markkj/hackathon-season2/internal/json"
	"github.com/markkj/hackathon-season2/internal/xml"
)

const (
	moveStatus = "1"
)

func main() {
	data, err := xml.ReadXMLFromHackathon("./data-devclub-1.xml")
	if err != nil {
		fmt.Println(err)
	}
	csvFile := &csv.CsvData{
		Columns: []string{},
		Records: []string{},
	}
	columnStrings := []string{
		"EMPID",
		"PASSPORT",
		"FIRSTNAME",
		"LASTNAME",
		"GENDER",
		"BIRTHDAY",
		"NATIONALITY",
		"HIRED",
		"DEPT",
		"POSITION",
		"STATUS",
		"REGION",
	}
	csvFile.SetColumn(columnStrings)
	for _, row := range data {
		record := make([]string, len(columnStrings))
		isValid := true
		for key, value := range row {
			if key == "HIRED" {
				date, _ := time.Parse("02-01-2006", value)
				date = time.Time{}.Add(time.Now().Sub(date))
				if date.Year() < 3 {
					isValid = false
				}
			}
			for i, c := range columnStrings {
				if key == c {
					record[i] = value
					break
				}
			}
		}
		if isValid {
			csvFile.AddRecord(record)
		}
	}
	path, _ := os.Getwd()
	path += "/q1.csv"
	err = csvFile.BuildCsvFile(path)
	if err != nil {
		fmt.Println(err)
	}
	records, err := csv.CSVFileToMap(path)
	if err != nil {
		fmt.Println(err)
	}
	json.ExportToJsonFile(records, "DevMountainAnwserQ2")
}
