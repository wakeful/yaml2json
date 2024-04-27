package parse

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func decode(input interface{}) interface{} {
	switch element := input.(type) {
	case map[interface{}]interface{}:
		rec := map[string]interface{}{}

		for key, value := range element {
			rec[fmt.Sprintf("%v", key)] = decode(value)
		}

		return rec
	case []interface{}:
		for key, value := range element {
			element[key] = decode(value)
		}
	}

	return input
}

func ByteSliceToJSON(input []byte) (string, error) {
	var unYaml interface{}
	if err := yaml.Unmarshal(input, &unYaml); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	output, err := json.Marshal(decode(unYaml))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return string(output), nil
}
