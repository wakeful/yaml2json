package main

import (
	"github.com/go-yaml/yaml"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	showVersion = flag.Bool("version", false, "show version and exit")
	url         = "https://github.com/wakeful/yaml2json"
	version     = "dev"
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("yaml2json\n url: %s\n version: %s\n", url, version)
		os.Exit(0)
	}

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
