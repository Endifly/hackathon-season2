package json

import (
	"encoding/json"
	"os"
)

func ExportToJsonFile(records []map[string]interface{}) error {
	csvFile, err := os.Create("DevMountainAnwser.json")
	if err != nil {
		return err
	}
	// adjust format of output
	csvFile.WriteString("[\n")
	defer csvFile.Close()
	for i, row := range records {
		jsonStr, err := json.MarshalIndent(row, "", " ")
		if err != nil {
			return err
		}
		//only add "," if not the last row
		str := string(jsonStr)
		if i != len(records)-1 {
			str = str + ","
		}
		_, err = csvFile.WriteString(str)
		if err != nil {
			return err
		}
	}
	// adjust format of output
	csvFile.WriteString("\n]")
	return nil
}
