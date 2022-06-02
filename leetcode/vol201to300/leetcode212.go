package main

func main() {

}

type TrieNode struct {
	s   string //保存的字符串
	son []*TrieNode
}

func (tr *TrieNode) insert(str string) {
	for i := range str {
		var u = str[i] - 'a'
		if tr.son[u] == nil {
			tr.son[u] = &TrieNode{s: "", son: make([]*TrieNode, 26)}
		}
		tr = tr.son[u]
	}
	tr.s = str
}

var moveX = []int{0, -1, 0, 1}
var moveY = []int{-1, 0, 1, 0}
var m = 0
var n = 0
var root TrieNode

func findWords(board [][]byte, words []string) (ret []string) {
	var rowLen = len(board)
	var colLen = len(board[0])

	m, n = rowLen, colLen
	visit := make([][]bool, rowLen)
	for i, _ := range visit {
		visit[i] = make([]bool, colLen)
	}

	root = TrieNode{s: "", son: make([]*TrieNode, 26)}
	strSet := map[string]int{}
	for _, v := range words {
		root.insert(v)
	}

	// find
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			u := board[i][j] - 'a'
			if root.son[u] != nil {
				visit[i][j] = true
				dfs(board, &visit, i, j, strSet, *root.son[u])
				visit[i][j] = false
			}
		}
	}

	for k, _ := range strSet {
		ret = append(ret, k)
	}

	return ret
}

func dfs(board [][]byte, visit *[][]bool, row int, col int, strSet map[string]int, node TrieNode) {
	if node.s != "" {
		strSet[node.s]++
	}

	for i := 0; i < 4; i++ {
		nowX := row + moveX[i]
		nowY := col + moveY[i]
		if isVaild(nowX, nowY) && (*visit)[nowX][nowY] == false {
			u := board[nowX][nowY] - 'a'
			if node.son[u] != nil {
				(*visit)[nowX][nowY] = true
				dfs(board, visit, nowX, nowY, strSet, *node.son[u])
				(*visit)[nowX][nowY] = false
			}
		}
	}

	return
}

func isVaild(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
