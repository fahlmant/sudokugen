package main

func main() {

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			sudokugen.fillNumber(i, j)
		}
	}

	sudokugen.printBoard()
}
