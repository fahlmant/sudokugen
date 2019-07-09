package sudokugen

import (
	"math/rand"
	"time"
)

//Coord represents a cell in the Board by (x,y)
type Coord struct {
	x, y int
}

//Cell represents a single value in the puzzle
type Cell struct {
	value int
}

//Box is on
type Box struct {
	//Some 2D array of cells
	cells interface{}
}

//Board represents the full puzzle board
type Board struct {
	//Some array of squares
	//Need to be broken up into rows as a function of dimension
	//grid []Box
	//length int
	//xDimension int
	//yDimension int
	//boxCount int

	//boxesInRow = length / xDimension
	//boxesInCol = length / yDimension
}

func inRow(num int, x int, y int) bool {

	var inrow = false

	return inrow
}

func inColumn(num int, x int, y int) bool {

	var incolumn = false

	return incolumn
}

func inBox(num int, x int, y int) bool {

	return false
}

func fillNumber(x int, y int) {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i <= 12; i++ {
		var num = r1.Intn(13-1) + 1
		if !inRow(num, x, y) && !inColumn(num, x, y) && !inBox(num, x, y) {
			//Fill cell
		}
	}
}

//GenerateBoard takes a number N and reutrns a fully solved NxN puzzle
func GenerateBoard(dimension int) Board {

	//Detrmine Box dimensions and number of Boxes
	//Generate as many as possible independent Boxes
	//For each other box, recursively fill
	return Board{}
}
