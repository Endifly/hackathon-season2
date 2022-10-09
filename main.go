package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
	employees, err := MapToEmployees(data)
	if err != nil {
		log.Fatal(err)
	}

	employees = employees.FilterBy(isValidStatus, isValidGender, isActiveStatus)
	fmt.Printf("clean result row = %d\n", len(employees))
	err = exportToCSV("clean", employees)
	if err != nil {
		log.Fatal(err)
	}

	// check role
	employees = employees.FilterBy(isOnlyPosition)
	fmt.Printf("check only positon steward,airhostess,pilot result row = %d\n", len(employees))
	err = exportToCSV("only_steward_airhostess_pilot_and_active", employees)
	if err != nil {
		log.Fatal(err)
	}

	// check exp more than three compare with current date
	employees = employees.FilterBy(isExpMoreThanThree)
	fmt.Printf("check only exp more than three result row = %d\n", len(employees))
	err = exportToCSV("only_exp_more_than_three", employees)
	if err != nil {
		log.Fatal(err)
	}

	err = exportToCSV("DevMountain", employees)
	if err != nil {
		log.Fatal(err)
	}

	groupByNation := employees.GroupByNation()
	for key, group := range groupByNation {
		err = exportToCSV(fmt.Sprintf("DevMountain-%s", key), group)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func exportToCSV(fileName string, employees Employees) error {
	csvFile := &csv.CsvData{
		Columns: columnStrings,
		Records: []string{},
	}
	for _, emp := range employees {
		record := make([]string, len(columnStrings))
		record[0] = strconv.Itoa(emp.EMPID)
		record[1] = emp.PASSPORT
		record[2] = emp.FIRSTNAME
		record[3] = emp.LASTNAME
		record[4] = strconv.Itoa(emp.GENDER)
		record[5] = emp.BIRTHDAY.Format("02-01-2006")
		record[6] = emp.NATIONALITY
		record[7] = emp.HIRED.Format("02-01-2006")
		record[8] = emp.DEPT
		record[9] = emp.POSITION
		record[10] = strconv.Itoa(emp.STATUS)
		record[11] = emp.REGION

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
