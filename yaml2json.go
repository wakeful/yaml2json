package main

import (
	"github.com/go-yaml/yaml"

	"encoding/json"
	"fmt"
	"log"
)

func main() {
	input, err := readInput()

	if err != nil {
		fmt.Println(err)
		fmt.Println()
		printHelp()
	}

	var content interface{}
	if err = yaml.Unmarshal([]byte(input), &content); err != nil {
		log.Fatalln(err)
	}

	if content, err := json.Marshal(decode(content)); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(string(content))
	}
}
