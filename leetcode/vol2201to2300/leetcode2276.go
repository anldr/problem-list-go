package main

func main() {

}

type CountIntervals struct {
	left, right *CountIntervals
	l, r, sum   int
}

func Constructor() CountIntervals { return CountIntervals{l: 1, r: 1e9} }

func (o *CountIntervals) fill() { o.sum = o.r - o.l + 1 } // 范围 [o.l,o.r] 内的所有整数都被区间覆盖

func (o *CountIntervals) Add(l, r int) {
	if l <= o.l && o.r <= r {
		o.fill() // lazy 更新：o 已被区间 [l,r] 包含，不再继续递归
		return
	}
	mid := (o.l + o.r) >> 1
	if o.left == nil {
		o.left = &CountIntervals{l: o.l, r: mid}
	} // 动态开点
	if o.right == nil {
		o.right = &CountIntervals{l: mid + 1, r: o.r}
	} // 动态开点
	if o.sum == o.r-o.l+1 { // o 被完整覆盖，那么左右节点也应当被完整覆盖
		o.left.fill()
		o.right.fill()
	}
	if l <= mid {
		o.left.Add(l, r)
	}
	if mid < r {
		o.right.Add(l, r)
	}
	o.sum = o.left.sum + o.right.sum
}

func (o *CountIntervals) Count() int {
	return o.sum
}
