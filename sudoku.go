//sudoku is a package that allows you to solve simple Sudoku puzzle in a 9x9 grid using a backtracking algorithm.
package sudoku

import (
	"errors"
	"fmt"
)

//Grid representing the 9x9 game grid.
type Grid [9][9]int

//Print allows the printing of the game grid to the stdout.
func (g Grid) Print() {
	for x := 0; x < len(g); x++ {
		fmt.Printf("[")
		for y := 0; y < len(g); y++ {
			fmt.Printf(" %v ", g[x][y])
		}
		fmt.Printf("]\n")
	}
}

//SolveGrid when given a Grid will try to solve the Sudoku puzzle.
//May return an error if no solution is found.
//If an error is returned, the returned grid is not valid.
//	var gameGrid Grid
//	gameGrid, err = sudoku.SolveGrid(gameGrid)
//	if err == nil {
//		gameGrid.Print()
//	}
func SolveGrid(grid Grid) (Grid, error) {
	var solvedGrid Grid
	success := backtracking(grid, &solvedGrid)

	if !success {
		return solvedGrid, errors.New("No solution found.")
	}

	return solvedGrid, nil
}

func backtracking(grid Grid, answer *Grid) bool {
	var row int = 0
	var col int = 0

	if !findUnassignedCell(grid, &row, &col) {
		*answer = grid
		return true
	}
	for num := 1; num <= 9; num++ {
		if isGridValid(grid, row, col, num) {
			grid[row][col] = num
			if backtracking(grid, answer) {
				return true
			} else {
				grid[row][col] = 0
			}
		}
	}
	return false
}

func isGridValid(grid Grid, row int, col int, num int) bool {
	return isRowValid(grid, row, num) &&
		isValidCol(grid, col, num) &&
		isSubGridValid(grid, row-row%3, col-col%3, num)
}

func isRowValid(grid Grid, row int, num int) bool {
	for col := 0; col < len(grid); col++ {
		if grid[row][col] == num {
			return false
		}
	}
	return true
}

func isValidCol(grid Grid, col int, num int) bool {
	for row := 0; row < len(grid); row++ {
		if grid[row][col] == num {
			return false
		}
	}
	return true
}

func isSubGridValid(grid Grid, boxStartRow int, boxStartCol int, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+boxStartRow][col+boxStartCol] == num {
				return false
			}
		}
	}
	return true
}

func findUnassignedCell(grid Grid, row *int, col *int) bool {
	for *row = 0; *row < len(grid); *row++ {
		for *col = 0; *col < len(grid); *col++ {
			if grid[*row][*col] == 0 {
				return true
			}
		}
	}
	return false
}
