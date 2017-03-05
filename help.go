package main

import (
	"fmt"
	"os"
)

const useText = `yaml2json usage example:

stdin pipe:
  cat file.yml | yaml2json

or specify a file:
  yaml2json path/file.yml
`

func printHelp() {
	fmt.Println(useText)
	os.Exit(1)
}
