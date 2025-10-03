package tools

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func Json2Csv(inputFile, outputFile string) error {
	// Open the input JSON file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	// Decode JSON data
	var data []map[string]string
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("no data to write")
	}

	// Extract headers from the first record
	headers := make([]string, 0, len(data[0]))
	for key := range data[0] {
		headers = append(headers, key)
	}
	sort.Strings(headers)

	// Create the output CSV file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Write CSV data
	writer := csv.NewWriter(outFile)

	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %v", err)
	}

	// Write each record
	for _, record := range data {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = record[header]
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write record: %v", err)
		}
	}

	// Flush the writer
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error flushing CSV writer: %v", err)
	}

	return nil
}
