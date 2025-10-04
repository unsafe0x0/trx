package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Json2Yaml(inputFile, outputFile string) error {
	jsonData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	yamlStr := toYAML(data, 0)

	if err := os.WriteFile(outputFile, []byte(yamlStr), 0644); err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}

	return nil
}

func toYAML(data interface{}, indent int) string {
	indentStr := strings.Repeat("  ", indent)

	switch v := data.(type) {
	case map[string]interface{}:
		var sb strings.Builder
		for key, val := range v {
			sb.WriteString(fmt.Sprintf("%s%s: %s\n", indentStr, key, scalarOrNested(val, indent+1)))
		}
		return sb.String()
	case []interface{}:
		var sb strings.Builder
		for _, item := range v {
			sb.WriteString(fmt.Sprintf("%s- %s\n", indentStr, scalarOrNested(item, indent+1)))
		}
		return sb.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func scalarOrNested(val interface{}, indent int) string {
	switch val.(type) {
	case map[string]interface{}, []interface{}:
		return "\n" + toYAML(val, indent)
	default:
		return fmt.Sprintf("%v", val)
	}
}
