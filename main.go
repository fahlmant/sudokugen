package main

import (
	"github.com/fahlmant/sudokugen/pkg/sudokugen"
)

func main() {

	board := sudokugen.GenerateBoard(4, 3)
	sudokugen.PrintBoard(board)
}
