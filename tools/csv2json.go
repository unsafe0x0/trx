package tools

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func Csv2Json(inputFile, outputFile string) error {
	// Open the input CSV file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	// Read CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}

	if len(records) == 0 {
		return fmt.Errorf("no data to write")
	}

	headers := records[0]
	var data []map[string]string

	for i, record := range records {
		if i == 0 { // Skip header row
			continue
		}
		m := make(map[string]string)
		for j, header := range headers {
			m[header] = record[j]
		}
		data = append(data, m)
	}

	// Create the output JSON file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Encode JSON data
	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ") // For pretty printing
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	return nil
}
