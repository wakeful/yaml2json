package main

func decode(input interface{}) interface{} {
	switch in := input.(type) {
	case map[interface{}]interface{}:
		rec := map[string]interface{}{}
		for k, v := range in {
			rec[k.(string)] = decode(v)
		}
		return rec
	case []interface{}:
		for k, v := range in {
			in[k] = decode(v)
		}
	}
	return input
}
