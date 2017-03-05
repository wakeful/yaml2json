package main

import (
	"fmt"
	"os"
)

func perr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
