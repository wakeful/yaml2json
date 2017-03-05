package main

import (
	"errors"
	"io/ioutil"
	"os"
)

func readInput() (input []byte, err error) {
	if stdinStat, _ := os.Stdin.Stat(); (stdinStat.Mode() & os.ModeCharDevice) == 0 {
		if input, err = ioutil.ReadAll(os.Stdin); err != nil {
			perr(err)
		}
	} else if len(os.Args) > 1 {
		if input, err = ioutil.ReadFile(os.Args[1]); err != nil {
			perr(err)
		}
	} else {
		return nil, errors.New("Missing input")
	}

	return
}
