package main

import (
	"os"
	"strconv"

	"github.com/fahlmant/sudokugen/pkg/sudokugen"
)

func validateInput(x string, y string) (int, int, bool) {

	xInt := 0
	yInt := 0

	xInt, err := strconv.Atoi(x)
	if err != nil {
		return 0, 0, false
	}
	yInt, err = strconv.Atoi(y)
	if err != nil {
		return 0, 0, false
	}

	return xInt, yInt, true
}

func main() {

	if len(os.Args[1:]) < 2 || len(os.Args[1:]) > 2 {
		sudokugen.PrintUsage()
		os.Exit(0)
	}

	x, y, isValid := validateInput(os.Args[1], os.Args[2])
	if isValid {
		board := sudokugen.GenerateBoard(x, y)
		sudokugen.PrintBoard(board)
	} else {
		sudokugen.PrintUsage()
	}
}
