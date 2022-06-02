package main

import (
	. "../common_utils"
	"fmt"
)

type SegeMentTree struct {
	val          int64
	lRng, rRng   int64
	lNode, rNode *SegeMentTree
}

func (tr *SegeMentTree) update(val int64) {
	tr.val += val
}

func (tr *SegeMentTree) pushUp() {
	lVal, rVal := int64(0), int64(0)
	if tr.lNode != nil {
		lVal = tr.lNode.val
	}
	if tr.rNode != nil {
		rVal = tr.rNode.val
	}
	tr.val = lVal + rVal
}

func (tr *SegeMentTree) pushDown() {
}

func (tr *SegeMentTree) modify(l, r int64, val int64) {
	lRng, rRng := tr.lRng, tr.rRng
	if l <= lRng && rRng <= r {
		tr.update(val)
	} else {
		mid := (lRng + rRng) >> 1
		if mid >= l {
			if tr.lNode == nil {
				tr.lNode = &SegeMentTree{val: 0, lRng: lRng, rRng: mid}
			}
			tr.lNode.modify(l, r, val)
		} else {
			if tr.rNode == nil {
				tr.rNode = &SegeMentTree{val: 0, lRng: mid + 1, rRng: rRng}
			}
			tr.rNode.modify(l, r, val)
		}
		tr.pushUp()
	}
}

func (tr *SegeMentTree) query(l, r int64) int {
	if tr == nil {
		return 0
	}

	L, R := tr.lRng, tr.rRng
	if l <= L && R <= r {
		return int(tr.val)
	} else {
		var ans = 0
		mid := (L + R) >> 1
		if mid >= l {
			ans += tr.lNode.query(l, r)
		}
		if r > mid {
			ans += tr.rNode.query(l, r)
		}
		return ans
	}
}

func countRangeSum(nums []int, lower int, upper int) int {
	lRng, rRng, preSum := int64(0), int64(0), int64(0)
	for _, v := range nums {
		preSum = preSum + int64(v)
		lRng = MinInt64(MinInt64(lRng, preSum), MinInt64(preSum-int64(upper), preSum-int64(lower)))
		rRng = MaxInt64(MaxInt64(rRng, preSum), MaxInt64(preSum-int64(upper), preSum-int64(lower)))
	}

	preSum = 0
	var ret = 0
	root := SegeMentTree{lRng: lRng, rRng: rRng, val: 0}
	root.modify(0, 0, 1)
	for _, v := range nums {
		preSum = preSum + int64(v)
		ret = ret + root.query(preSum-int64(upper), preSum-int64(lower))
		root.modify(preSum, preSum, 1)
	}

	return ret
}

func main() {
	fmt.Println(countRangeSum([]int{2147483647, -2147483648, -1, 0}, -1, 0))
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
