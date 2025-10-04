package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Yaml2Json(inputFile, outputFile string) error {
	yamlData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	data, err := parseYAML(string(yamlData))
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %v", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	return nil
}

func parseYAML(yamlStr string) (interface{}, error) {
	lines := strings.Split(yamlStr, "\n")
	return parseLines(lines, 0, 0)
}

func parseLines(lines []string, startIndent, index int) (interface{}, error) {
	var result interface{}
	mapResult := make(map[string]interface{})
	arrayResult := []interface{}{}

	for i := index; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}

		indent := countIndent(line)
		if indent < startIndent {
			break
		}

		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "- ") {
			item := strings.TrimSpace(trimmed[2:])
			if item == "" {
				val, err := parseLines(lines, indent+2, i+1)
				if err != nil {
					return nil, err
				}
				arrayResult = append(arrayResult, val)
			} else {
				arrayResult = append(arrayResult, item)
			}
			result = arrayResult
		} else if strings.Contains(trimmed, ":") {
			parts := strings.SplitN(trimmed, ":", 2)
			key := strings.TrimSpace(parts[0])
			valStr := strings.TrimSpace(parts[1])
			if valStr == "" {
				val, err := parseLines(lines, indent+2, i+1)
				if err != nil {
					return nil, err
				}
				mapResult[key] = val
			} else {
				mapResult[key] = valStr
			}
			result = mapResult
		}
	}

	return result, nil
}

func countIndent(line string) int {
	count := 0
	for _, ch := range line {
		if ch == ' ' {
			count++
		} else {
			break
		}
	}
	return count
}
