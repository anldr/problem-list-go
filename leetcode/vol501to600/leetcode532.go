package main

import (
	. "../common_utils"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 0))
	fmt.Println(findPairs([]int{1, 1, 1, 1, 1}, 0))
	fmt.Println(findPairs([]int{1, 1, 1, 2, 2}, 0))
}

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	result := 0
	arrLen := len(nums)
	left, right := 0, 1
	for right < arrLen && left < right {
		if AbsInt32(nums[right]-nums[left]) < k {
			right++
		} else if AbsInt32(nums[right]-nums[left]) > k {
			left++
		} else {
			result++
			pos := left
			for ; pos < arrLen && nums[left] == nums[pos]; pos++ {
			}
			left = pos
			pos = right
			for ; pos < arrLen && nums[right] == nums[pos]; pos++ {
			}
			right = pos

		}
		if left == right {
			right++
		}
	}

	return result
}
