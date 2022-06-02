package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasValidPath([][]byte{{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'}}))
}

func hasValidPath(grid [][]byte) bool {
	if grid[0][0] == ')' {
		return false
	}
	row, col := len(grid), len(grid[0])

	if (row+col-1)&1 == 1 {
		return false
	}

	return dfs(grid, row, col, 0, 0, 1, 0)
}

func dfs(grid [][]byte, row, col int, x, y int, lp, rp int) bool {
	if x == (row-1) && y == (col-1) {
		return lp == rp
	}

	var nx = x
	var ny = y + 1
	if isValid(nx, ny, row, col) {
		if grid[nx][ny] == '(' {
			if (lp + 1) <= (row+col-1)/2 {
				if dfs(grid, row, col, nx, ny, lp+1, rp) {
					return true
				}
			}
		} else {
			if (rp+1) <= lp && (rp+1) <= (row+col-1)/2 {
				if dfs(grid, row, col, nx, ny, lp, rp+1) {
					return true
				}
			}
		}
	}

	nx = x + 1
	ny = y
	if isValid(nx, ny, row, col) {
		if grid[nx][ny] == '(' {
			if (lp + 1) <= (row+col-1)/2 {
				if dfs(grid, row, col, nx, ny, lp+1, rp) {
					return true
				}
			}
		} else {
			if (rp+1) <= lp && (rp+1) <= (row+col-1)/2 {
				if dfs(grid, row, col, nx, ny, lp, rp+1) {
					return true
				}
			}
		}
	}

	return false
}

func isValid(x, y int, row, col int) bool {
	return x >= 0 && x < row && y >= 0 && y < col
}
