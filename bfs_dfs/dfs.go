// https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	maxY := len(board)
	maxX := len(board[0])

	// loop for four edges
	// if it is an 'O', do dfs to mark all connected 'O' to '#'
	// because all connected 'O' with bondary edgs are not counted
	// so we can flip all remain 'O' to 'X'
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if y == 0 || y == maxY-1 || x == 0 || x == maxX-1 {
				if board[y][x] == 'O' {
					dfs(board, y, x)
				}
			}
		}
	}

	// filp all '#' to 'O' and remain 'O' to 'X'
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, y, x int) {
	if y < 0 || x < 0 || y >= len(board) || x >= len(board[0]) {
		return
	}
	if board[y][x] == 'O' {
		board[y][x] = '#'
		dfs(board, y-1, x)
		dfs(board, y+1, x)
		dfs(board, y, x-1)
		dfs(board, y, x+1)
	}
}