package main

import "github.com/fahlmant/sudokugen/pkg/sudokugen"

func main() {

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			sudokugen.FillNumber(i, j)
		}
	}

	sudokugen.PrintBoard()
}
