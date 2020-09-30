package main

import "fmt"

func main() {
	grid := [][]int{
		{7, 0, 5},
		{2, 4, 6},
		{3, 8, 1},
	}

	fmt.Println(numMagicSquaresInside(grid))
}

func numMagicSquaresInside(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	count := 0

	for i := 0; i < row-2; i++ {
		for j := 0; j < col-2; j++ {

			if !valid(grid, i, j) {
				break
			}

			magic := true

			hsum := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			vsum := grid[i][j] + grid[i+1][j] + grid[i+2][j]
			dsum := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]

			for y := 1; y <= 2; y++ {
				if hsum != horiSum(grid, i, j, y) {
					magic = false
					break
				}
			}
			if !magic {
				break
			}

			for x := 1; x <= 2; x++ {
				if vsum != vertSum(grid, i, j, x) {
					magic = false
					break
				}
			}
			if !magic {
				break
			}

			if dsum != diagSum(grid, i, j, 2) {
				magic = false
				break
			}
			if !magic {
				break
			}

			count++

		}
	}

	return count
}

func vertSum(grid [][]int, i, j, x int) int {
	return grid[i][j+x] + grid[i+1][j+x] + grid[i+2][j+x]
}

func horiSum(grid [][]int, i, j, y int) int {
	return grid[i+y][j] + grid[i+y][j+1] + grid[i+y][j+2]
}

func diagSum(grid [][]int, i, j, k int) int {
	return grid[i][j+k] + grid[i+1][j+1] + grid[i+2][j+2-k]
}

func valid(grid [][]int, i, j int) bool {
	elem := grid[i][j]
	distinct := false
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if grid[i+x][j+y] > 9 || grid[i+x][j+y] < 1 {
				return false
			}
			if elem != grid[i+x][j+y] {
				distinct = true
			}
		}
	}
	return distinct
}
