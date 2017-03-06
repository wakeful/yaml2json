package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func readInput() (input []byte, err error) {
	if stdinStat, _ := os.Stdin.Stat(); (stdinStat.Mode() & os.ModeCharDevice) == 0 {
		if input, err = ioutil.ReadAll(os.Stdin); err != nil {
			log.Fatalln(err)
		}
	} else if len(os.Args) > 1 {
		if input, err = ioutil.ReadFile(os.Args[1]); err != nil {
			log.Fatalln(err)
		}
	} else {
		err = errors.New("Missing input")
	}

	return
}
