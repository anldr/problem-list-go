package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{10, 0}, 2))
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
}

func smallestRangeI(nums []int, k int) int {
	minn, maxx := 999999, -1
	for _, v := range nums {
		minn = Min(v, minn)
		maxx = Max(v, maxx)
	}

	if minn+k >= maxx-k {
		return 0
	}

	return maxx - 2*k - minn
}

func Min(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

func Max(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}
