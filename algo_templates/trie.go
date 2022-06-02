package algo_templates

type TrieNode struct {
	num   int
	child []*TrieNode
}

/* 插入一个字符串 */
func (tr *TrieNode) Insert(str string) {
	for i := range str {
		var u = str[i] - 'a'
		if tr.child[u] == nil {
			tr.child[u] = &TrieNode{num: 0, child: make([]*TrieNode, 26)}
		}
		tr = tr.child[u]
	}
	tr.num++
}

/* 查询字符串出现次数 */
func (tr *TrieNode) Query(str string) int {
	for i := range str {
		var u = str[i] - 'a'
		if tr.child[u] == nil {
			return 0
		}
		tr = tr.child[u]
	}
	return tr.num
}
