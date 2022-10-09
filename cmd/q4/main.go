package main

import (
	"fmt"

	"github.com/markkj/hackathon-season2/internal/csv"
	"github.com/markkj/hackathon-season2/internal/xml"
)

const ()

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
	err = csvFile.BuildCsvFile("../../q4.csv")
	if err != nil {
		fmt.Println(err)
	}

}
