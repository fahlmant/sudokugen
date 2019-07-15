package sudokugen

import (
	"fmt"
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

//Box is
type Box struct {
	//Some 2D array of cells
	cells interface{}
}

//Board represents the full puzzle board
type Board struct {
	//Horizontal dimension of one Box
	xDimension int
	//Vertial dimension of one Box
	yDimension int
	//The full grid
	grid [][]int
	//Length of one full size of the board
	fullDimension int
	//Number of boxes in a row
	boxesInRow int
	//Number if boxes in a column
	boxesInCol int
}

//NewBoard creates a new board based on the dimensions of a single Box
func NewBoard(x int, y int) Board {

	dimension := x * y
	grid := make([][]int, dimension)
	for i := range grid {
		grid[i] = make([]int, dimension)
	}

	//Create new board with x and y dim on one box, total dimension,
	//and number of boxes in a row and column
	return Board{x, y, grid, dimension, y, x}
}

//Checks if value is in the row
func (b *Board) inRow(c Coord, num int) bool {

	inrow := false
	for row := 0; row < b.fullDimension; row++ {
		if row != c.x {
			//fmt.Printf("Row check: Index %v, Cell: %v, Value: %v\n", row, b.getValue(Coord{row, c.y}), num)
			if num == b.getValue(Coord{row, c.y}) {
				inrow = true
			}
		}
	}

	//fmt.Printf("In row: %v\n", inrow)
	return inrow
}

//Checks if value is in the column
func (b *Board) inColumn(c Coord, num int) bool {

	incolumn := false
	for col := 0; col < b.fullDimension; col++ {
		if col != c.y {
			//fmt.Printf("Row check: Index %v, Cell: %v, Value: %v\n", col, b.getValue(Coord{c.x, col}), num)
			if num == b.getValue(Coord{c.x, col}) {
				incolumn = true
			}
		}
	}

	//fmt.Printf("In col: %v\n", incolumn)
	return incolumn
}

func (b *Board) inBox(c Coord, num int) bool {

	inbox := false

	//fmt.Printf("In box: %v\n", inbox)
	return inbox
}

//Gets the index of a box by the current cell index
//Box idex starts at 1, moving horizontally for each verical line
//I.e, in a 9x9 sudoku, the top row would be Boxes 1, 2 and 3 from left to right
func (b *Board) getBoxIndex(currentCell int) int {

	for i := 1; i <= (b.fullDimension); i++ {
		if currentCell <= ((b.xDimension * b.yDimension) * i) {
			return i
		}
	}
	return 0
}

func (b *Board) getValue(c Coord) int {

	return b.grid[c.y][c.x]
}

func (b *Board) findEmptyCell() (bool, Coord) {

	for y := 0; y < b.fullDimension; y++ {
		for x := 0; x < b.fullDimension; x++ {
			c := Coord{x, y}
			if b.getValue(c) == 0 {
				//fmt.Printf("Index: %v, Value: %v", c, b.getValue(c))
				//fmt.Println("Found empty cell")
				return true, c
			}
		}
	}
	return false, Coord{}
}

func (b *Board) isSafe(c Coord, num int) bool {

	if b.inColumn(c, num) || b.inRow(c, num) || b.inBox(c, num) {
		return false
	}

	return true
}

func (b *Board) fillCell(c Coord, num int) {
	b.grid[c.y][c.x] = num
}

func (b *Board) getRandBoardValue() int {

	time.Sleep(500)
	num := (rand.Intn((b.fullDimension + 1) - 1)) + 1
	return num
}

func (b *Board) fillBoard() bool {
	//bactracking algorithm
	//1. Find empty cell
	//2. Tentatively fill a number
	//3. Recursive call
	//4. If recursive call returns false, set cell empty

	//fmt.Println(b.grid)
	//Find an empty cell
	needsValue, cell := b.findEmptyCell()

	//If there are no empty cells, board is complete
	if !needsValue {
		return true
	}

	for i := 0; i < b.fullDimension; i++ {
		//Get a random value between 1 and N
		value := b.getRandBoardValue()
		//If it's safe to put the value in the open cell
		if b.isSafe(cell, value) {
			//Tentatively place value in the cell
			b.fillCell(cell, value)
			//If this value allows for valid solution, return true
			if b.fillBoard() {
				return true
			}
			//If the value is not part of a valid solution, reset it to empty
			b.fillCell(cell, 0)
		}

	}

	//Triggers backtrack
	return false
}

//GenerateBoard takes a number N and reutrns a fully solved NxN puzzle
func GenerateBoard(x int, y int) Board {

	rand.Seed(time.Now().UnixNano())
	board := NewBoard(x, y)
	board.fillBoard()
	fmt.Println(board.grid)

	return board
}
