package parse

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadInput() ([]byte, error) {
	if len(os.Args) > 1 {
		input, err := os.ReadFile(os.Args[1])
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return input, nil
	}

	if stdinStat, _ := os.Stdin.Stat(); (stdinStat.Mode() & os.ModeCharDevice) == 0 {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return input, nil
	}

	return nil, errMissingInput
}

var errMissingInput = errors.New("missing input")
