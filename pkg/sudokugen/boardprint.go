package sudokugen

import "fmt"

func PrintRowOfBoxes(board Board, index int) {

	boxIndex := index - 1
	for y := board.yDimension * boxIndex; y < board.yDimension*index; y++ {
		for x := 0; x < board.fullDimension; x++ {
			fmt.Printf("%v ", board.grid[y][x])
			if ((x + 1) % board.xDimension) == 0 {
				fmt.Printf("    ")
			}
		}
		fmt.Println("")
	}

}

func PrintColumnOfBoxes(board Board, index int) {

	boxIndex := index - 1
	for y := 0; y < board.fullDimension; y++ {
		for x := board.xDimension * boxIndex; x < board.xDimension*index; x++ {
			fmt.Printf("%v ", board.grid[y][x])
		}
		fmt.Println("")
		if ((y + 1) % board.yDimension) == 0 {
			fmt.Println("")
		}
	}
}

func printBox(board Board, index int) {

}

func PrintBoard(board Board) {

	for row := 1; row <= board.yDimension; row++ {
		PrintRowOfBoxes(board, row)
		fmt.Println("")
	}
}
