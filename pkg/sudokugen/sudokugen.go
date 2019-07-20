package sudokugen

import (
	"math"
	"math/rand"
	"time"
)

//Coord represents a cell in the Board by (x,y)
type Coord struct {
	x, y int
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

	xBoxIndex, yBoxIndex := b.getBoxIndex(c)
	for x := (xBoxIndex * b.xDimension); x < (xBoxIndex*b.xDimension)+b.xDimension; x++ {
		for y := (yBoxIndex * b.yDimension); y < (yBoxIndex*b.yDimension)+b.yDimension; y++ {
			if num == b.getValue(Coord{x, y}) {
				inbox = true
			}
		}
	}
	return inbox
}

//Gets the index of a box by the current cell index
//Box idex starts at 0, moving horizontally for each verical line
//I.e, in a 9x9 sudoku, the top row would be Boxes 0, 1 and 2 from left to right
func (b *Board) getBoxIndex(currentCell Coord) (xBoardIndex int, yBoardIndex int) {

	//Calculate x index of Box
	xBoxIndex := math.Floor(float64(currentCell.x / b.xDimension))
	//Calculate y index of Box
	yBoxIndex := math.Floor(float64(currentCell.y / b.yDimension))

	return int(xBoxIndex), int(yBoxIndex)
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

	//time.Sleep(500)
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

	return board
}
