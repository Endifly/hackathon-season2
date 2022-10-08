package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/markkj/hackathon-season2/internal/csv"
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
		for key, value := range row {
			if key == "STATUS" {
				if value != moveStatus {
					continue
				}
			}
			if key == "POSITION" {
				if !strings.Contains("Airhostess,Pilot,Steward", value) {
					continue
				}
			}
			if key == "HIRED" {
				date, _ := time.Parse("02-01-2006", value)
				fmt.Println(time.Now().Sub(date))
			}
			for i, c := range columnStrings {
				if key == c {
					record[i] = value
					break
				}
			}
		}
		csvFile.AddRecord(record)
	}
	err = csvFile.BuildCsvFile()
	if err != nil {
		fmt.Println(err)
	}
}
