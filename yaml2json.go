package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
)

func main() {
	input, err := readInput()

	if err != nil {
		printHelp()
	}

	var content interface{}
	if err = yaml.Unmarshal([]byte(input), &content); err != nil {
		perr(err)
	}

	if content, err := json.Marshal(decode(content)); err != nil {
		perr(err)
	} else {
		fmt.Println(string(content))
	}
}
