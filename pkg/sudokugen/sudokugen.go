package sudokugen

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	//The board
	grid [12][12]int

	//There's probably a better programatic way to do this
	//Top Row
	square1Top    = Coord{0, 0}
	square1Bottom = Coord{3, 2}
	square1       = Square{square1Top, square1Bottom}

	square2Top    = Coord{4, 0}
	square2Bottom = Coord{7, 2}
	square2       = Square{square2Top, square2Bottom}

	square3Top    = Coord{8, 0}
	square3Bottom = Coord{11, 2}
	square3       = Square{square3Top, square3Bottom}

	//Middle Top Row
	square4Top    = Coord{0, 3}
	square4Bottom = Coord{3, 5}
	square4       = Square{square4Top, square4Bottom}

	square5Top    = Coord{4, 3}
	square5Bottom = Coord{7, 5}
	square5       = Square{square5Top, square5Bottom}

	square6Top    = Coord{8, 3}
	square6Bottom = Coord{11, 5}
	square6       = Square{square6Top, square6Bottom}

	//Middle Bottom Row
	square7Top    = Coord{0, 6}
	square7Bottom = Coord{3, 8}
	square7       = Square{square7Top, square7Bottom}

	square8Top    = Coord{4, 6}
	square8Bottom = Coord{7, 8}
	square8       = Square{square8Top, square8Bottom}

	square9Top    = Coord{8, 6}
	square9Bottom = Coord{11, 8}
	square9       = Square{square9Top, square9Bottom}

	//Bottom Row
	square10Top    = Coord{0, 9}
	square10Bottom = Coord{3, 11}
	square10       = Square{square10Top, square10Bottom}

	square11Top    = Coord{4, 9}
	square11Bottom = Coord{7, 11}
	square11       = Square{square11Top, square11Bottom}

	square12Top    = Coord{8, 9}
	square12Bottom = Coord{11, 11}
	square12       = Square{square12Top, square12Bottom}
)

type Coord struct {
	x, y int
}

type Square struct {
	topLeft     Coord
	bottomRight Coord
}

func inRow(num int, x int, y int) bool {

	var inrow = false

	for i := 0; i < 12; i++ {
		if i != y {
			if num == grid[x][i] {
				inrow = true
			}
		}
	}

	return inrow
}

func inColumn(num int, x int, y int) bool {

	var incolumn = false

	for i := 0; i < 12; i++ {
		if i != x {
			if num == grid[i][y] {
				incolumn = true
			}
		}
	}

	return incolumn
}

func inSquare(num int, x int, y int) bool {

	return false
}

func FillNumber(x int, y int) {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i <= 12; i++ {
		var num = r1.Intn(13-1) + 1
		if !inRow(num, x, y) && !inColumn(num, x, y) && !inSquare(num, x, y) {
			grid[x][y] = num
		}
	}
}

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
