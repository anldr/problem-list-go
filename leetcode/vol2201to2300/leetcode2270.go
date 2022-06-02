package main

import "fmt"

func main() {
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
	fmt.Println(waysToSplitArray([]int{2, 3, 1, 0}))
}

func waysToSplitArray(nums []int) int {
	var arrLen = len(nums)
	var preSum = make([]int, arrLen+1)

	preSum[0] = 0
	for i := 1; i <= arrLen; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	var ret = 0
	for i := 0; i < arrLen-1; i++ {
		if preSum[i+1] >= (preSum[arrLen] - preSum[i+1]) {
			ret++
		}
	}

	return ret
}
