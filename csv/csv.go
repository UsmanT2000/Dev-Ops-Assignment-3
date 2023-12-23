package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func isCSVEmpty(records []string) bool {
	for _, record := range records {
		if record != "" {
			return false
		}
	}
	return true
}

func ParseCSV(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var csvData []map[string]string
	header := records[0] //ignoring column names
	if len(header) == 0 {
		return csvData, fmt.Errorf("Csv missing column Names")
	}
	for _, record := range records[1:] { //start iterating from where data values are(first row)
		if isCSVEmpty(record) {
			return csvData, fmt.Errorf("Csv contains empty records")
		}
		row := make(map[string]string)
		for i, value := range record {
			if i < len(header) {
				row[header[i]] = value
			}
		}
		csvData = append(csvData, row)
	}

	return csvData, nil
}
