package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestDistancePair([]int{1, 6, 1}, 3))
	fmt.Println(smallestDistancePair([]int{1, 1, 1}, 2))
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 1))
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, 1000000
	for left < right {
		mid := (left + right) >> 1
		if checkFun(mid, nums, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func checkFun(val int, nums []int, k int) bool {
	result := 0
	arrLen := len(nums)
	for i, j := 0, 1; i < arrLen; i++ {
		for ; j < arrLen && nums[j]-nums[i] <= val; j++ {
		}
		result = result + j - i - 1
	}
	return result >= k
}
