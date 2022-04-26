package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'C', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'C', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'C', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}

var moveX = []int{0, -1, 0, 1}
var moveY = []int{-1, 0, 1, 0}
var m = 0
var n = 0
var strLen = 0

func exist(board [][]byte, word string) bool {
	var rowLen = len(board)
	var colLen = len(board[0])

	strLen = len(word)
	m, n = rowLen, colLen
	visit := make([][]bool, rowLen)
	for i, _ := range visit {
		visit[i] = make([]bool, colLen)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visit[i][j] = true
				if dfs79(board, &visit, i, j, 1, word) {
					return true
				}
				visit[i][j] = false
			}
		}
	}
	return false
}

func dfs79(board [][]byte, visit *[][]bool, row int, col int, deep int, word string) bool {
	if deep >= strLen {
		return true
	}

	for i := 0; i < 4; i++ {
		nowX := row + moveX[i]
		nowY := col + moveY[i]
		if isVaild(nowX, nowY) && (*visit)[nowX][nowY] == false && board[nowX][nowY] == word[deep] {
			(*visit)[nowX][nowY] = true
			if dfs79(board, visit, nowX, nowY, deep+1, word) {
				return true
			}
			(*visit)[nowX][nowY] = false
		}
	}

	return false
}

func isVaild(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
