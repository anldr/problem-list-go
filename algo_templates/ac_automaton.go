package algo_templates

import "fmt"

type AcNode struct {
	len          int
	data         byte
	failPtr      *AcNode
	child        []*AcNode
	isEndingChar bool
}

func (ac *AcNode) insert(word string) {
	for _, v := range word {
		u := int(v - 'a')
		if ac.child[u] == nil {
			ac.child[u] = &AcNode{
				len: -1, data: byte(v), isEndingChar: false,
				failPtr: nil, child: make([]*AcNode, 26),
			}
		}
		ac = ac.child[u]
	}
	ac.len = len(word)
	ac.isEndingChar = true
}

func (ac *AcNode) match(text string) {
	p := ac
	for i, v := range text {
		u := int(v - 'a')
		for p.child[u] == nil && p != ac {
			p = p.failPtr
		}
		p = p.child[u]
		if p == nil {
			p = ac
		}
		temp := p
		for temp != ac {
			if temp.isEndingChar {
				pos := i - temp.len + 1
				fmt.Println("Begin Idx: %d , lenght: %d", pos, temp.len)
			}
			temp = temp.failPtr
		}
	}
}

func (ac *AcNode) buildFailPtr() {
	var queue = []AcNode{}
	ac.failPtr = nil
	queue = append(queue, *ac)
	for len(queue) > 0 {
		parent := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for i := 0; i < 26; i++ {
			pChild := parent.child[i]
			if pChild == nil {
				continue
			}
			if &parent == ac {
				pChild.failPtr = ac
			} else {
				qNode := parent.failPtr
				for qNode != nil {
					qc := qNode.child[pChild.data-'a']
					if qc != nil {
						pChild.failPtr = qc
						break
					}
					qNode = qNode.failPtr
				}
				if qNode == nil {
					pChild.failPtr = ac
				}
			}
			queue = append(queue, *pChild)
		}
	}
}
