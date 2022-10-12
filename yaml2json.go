package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wakeful/yaml2json/pkg/parse"
)

var (
	showVersion = flag.Bool("version", false, "show version and exit")
	version     = "dev"
)

func main() {
	flag.Parse()

	const url = "https://github.com/wakeful/yaml2json"
	if *showVersion {
		fmt.Printf("yaml2json\n url: %s\n version: %s\n", url, version)
		os.Exit(0)
	}

	input, err := parse.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	output, err := parse.ByteSliceToJSON(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
