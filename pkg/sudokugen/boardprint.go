package sudokugen

import "fmt"

func printRowOfSquares(topLeft Coord, bottomRight Coord) {

	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			fmt.Print(grid[y][x])
			fmt.Print(" ")
			if (x+1)%4 == 0 {
				fmt.Print(" ")
				fmt.Print(" ")
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func printColumnOfSquares(topLeft Coord, bottomRight Coord) {

}

func printSquare(top Coord, bottom Coord) {

	for y := top.y; y <= bottom.y; y++ {
		for x := top.x; x <= bottom.x; x++ {
			fmt.Print(grid[y][x])
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func PrintBoard() {

	for x := 0; x < 12; x++ {
		fmt.Println(grid[x])
	}
	fmt.Println("")
	printRowOfSquares(square1Top, square3Bottom)
	printRowOfSquares(square4Top, square6Bottom)
	printRowOfSquares(square7Top, square9Bottom)
	printRowOfSquares(square10Top, square12Bottom)
}
