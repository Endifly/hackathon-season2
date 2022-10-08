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
	defer csvFile.Close()
	for _, row := range records {
		jsonStr, err := json.Marshal(row)
		if err != nil {
			return err
		}
		_, err = csvFile.WriteString(string(jsonStr))
		if err != nil {
			return err
		}
	}
	return nil
}
