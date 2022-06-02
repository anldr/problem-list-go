package algo_templates

type SegeMentTree struct {
	lRng, rRng   int
	sum, lazy    int64
	lNode, rNode *SegeMentTree
}

func (tr *SegeMentTree) update(x int64) {
	tr.sum = tr.sum + int64(tr.rRng-tr.lRng+1)*x
	tr.lazy = tr.lazy + x
}

func (tr *SegeMentTree) pushUp() {
	tr.sum = tr.sum + tr.lNode.sum + tr.rNode.sum
}

func (tr *SegeMentTree) pushDown() {
	lazyVal := tr.lazy
	if lazyVal != 0 {
		tr.lNode.update(lazyVal)
		tr.rNode.update(lazyVal)
	}
	tr.lazy = 0
}

func (tr *SegeMentTree) build(l, r int, arr *[]int) {
	tr.lRng, tr.rRng = l, r
	tr.sum, tr.lazy = 0, 0
	if l == r {
		tr.sum = int64((*arr)[l])
	} else {
		mid := (l + r) >> 1
		tr.lNode = &SegeMentTree{}
		tr.lNode.build(l, mid, arr)
		tr.rNode = &SegeMentTree{}
		tr.rNode.build(mid+1, r, arr)
		tr.pushUp()
	}
}

func (tr *SegeMentTree) modify(l, r int, val int64) {
	L, R := tr.lRng, tr.rRng
	if l <= L && R <= r {
		tr.update(val)
	} else {
		tr.pushDown()
		mid := (L + R) >> 1
		if mid >= l {
			tr.lNode.modify(l, r, val)
		}
		if r > mid {
			tr.rNode.modify(l, r, val)
		}
		tr.pushUp()
	}
}

func (tr *SegeMentTree) query(l, r int) int64 {
	L, R := tr.lRng, tr.rRng
	if l <= L && R <= r {
		return tr.sum
	} else {
		tr.pushDown()
		var ans = int64(0)
		mid := (L + R) >> 1
		if mid >= l {
			ans = ans + tr.lNode.query(l, r)
		}
		if r > mid {
			ans = ans + tr.rNode.query(l, r)
		}
		tr.pushUp()
		return ans
	}
}
