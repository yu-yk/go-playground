package main

import (
	"fmt"
)

var entries = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func main() {
	board := [][]rune{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	printBoard(board)
	fmt.Println("Press enter to continue:")
	fmt.Scanln()
	var t bool = false
	finish := &t
	solveSudoku(board, finish)
	printBoard(board)
}

func solveSudoku(board [][]rune, finish *bool) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == '.' {
				for k := 0; k < 9; k++ {
					n := entries[k]
					if possible(x, y, n, board) {
						board[x][y] = n
						solveSudoku(board, finish)
						if !*finish {
							board[x][y] = '.'
						}
					}
				}
				return
			}
		}
	}
	*finish = true
	return
}

func possible(x, y int, n rune, board [][]rune) bool {
	// check horizontal
	for k := 0; k < 9; k++ {
		if board[x][k] == n {
			return false
		}
	}
	// check veritical
	for k := 0; k < 9; k++ {
		if board[k][y] == n {
			return false
		}
	}
	// check small square
	// sqX, sqY = 0/1/2
	sqX := x / 3
	sqY := y / 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[3*sqX+i][3*sqY+j] == n {
				return false
			}
		}
	}

	return true
}

func printBoard(board [][]rune) {
	for r := 0; r < len(board); r++ {
		fmt.Printf("%c\n", board[r])
		// for c := 0; c < len(board[r]); c++ {
		// 	fmt.Printf("%c")
		// }
	}
}
