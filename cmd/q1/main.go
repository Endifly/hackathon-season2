package main

import (
	"fmt"
	"log"
	"os"

	"github.com/markkj/hackathon-season2/internal/csv"
	"github.com/markkj/hackathon-season2/internal/json"
	"github.com/markkj/hackathon-season2/internal/xml"
)

const (
	moveStatus = "1"
)

var (
	columnStrings = []string{
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
)

func main() {
	data, err := xml.ReadXMLFromHackathon("./data-devclub-1.xml")
	if err != nil {
		log.Fatal(err)
	}
	data = filterBy(data, isValidStatus, isValidGender)

	err = exportToCSV("clean", data)
	if err != nil {
		log.Fatal(err)
	}

	data = filterBy(data, isActiveStatus, isOnlyPosition)
	if err != nil {
		log.Fatal(err)
	}

}

func exportToCSV(fileName string, data []map[string]string) error {
	csvFile := &csv.CsvData{
		Columns: columnStrings,
		Records: []string{},
	}
	for _, row := range data {
		record := make([]string, len(columnStrings))
		// if row["STATUS"] != moveStatus {
		// 	continue
		// }
		// if !strings.Contains("Airhostess,Pilot,Steward", row["POSITION"]) {
		// 	continue
		// }
		// date, _ := time.Parse("02-01-2006", row["HIRED"])
		// date = time.Time{}.Add(time.Now().Sub(date))
		// if date.Year() < 3 {
		// 	continue
		// }

		for key, value := range row {
			for i, c := range columnStrings {
				if key == c {
					record[i] = value
					break
				}
			}
		}
		csvFile.AddRecord(record)
	}
	path, _ := os.Getwd()
	path += fmt.Sprintf("/%s.csv", fileName)
	err := csvFile.BuildCsvFile(path)
	if err != nil {
		log.Fatal(err)
	}
	records, err := csv.CSVFileToMap(path)
	if err != nil {
		log.Fatal(err)
	}
	json.ExportToJsonFile(records, fileName)
	return nil
}

// FilterOption is a type function
type FilterOption func(data map[string]string) bool

func filterBy(data []map[string]string, filterOption ...FilterOption) []map[string]string {
	for i := 0; i < len(data); i++ {
		row := data[i]
		isValid := true

		for _, filter := range filterOption {
			if !filter(row) {
				isValid = false
				break
			}
		}

		if !isValid {
			data = append(data[:i], data[i+1:]...)
			i = i - 1
		}
	}

	return data
}

func isOnlyPosition(data map[string]string) bool {
	positions := []string{"Airhostess", "Pilot", "Steward"}
	for _, position := range positions {
		if position == data["POSITION"] {
			return true
		}
	}
	return false
}

func isActiveStatus(data map[string]string) bool {
	status := data["STATUS"]
	if status != "1" {
		return false
	}
	return true
}

func isValidStatus(data map[string]string) bool {
	status := data["STATUS"]
	if status != "1" && status != "2" && status != "3" {
		return false
	}
	return true
}

func isValidGender(data map[string]string) bool {
	gender := data["STATUS"]
	if gender != "0" && gender != "1" {
		return false
	}
	return true
}
