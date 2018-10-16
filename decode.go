package main

import (
	"fmt"
	"strconv"
)

func decode(input interface{}) interface{} {
	switch in := input.(type) {
	case map[interface{}]interface{}:
		rec := map[string]interface{}{}
		for k, v := range in {

			switch k.(type) {
			case bool:
				rec[strconv.FormatBool(k.(bool))] = decode(v)
			case int:
				rec[strconv.Itoa(k.(int))] = decode(v)
			case float64:
				rec[fmt.Sprintf("%f", k.(float64))] = decode(v)
			case string:
				rec[k.(string)] = decode(v)
			}

		}
		return rec
	case []interface{}:
		for k, v := range in {
			in[k] = decode(v)
		}
	}
	return input
}
