package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	var arrLen = len(nums)
	var result = make([]int, arrLen)

	result[0] = 1
	for i := 1; i < arrLen; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	var temp = 1
	for i := arrLen - 1; i >= 0; i-- {
		result[i] = result[i] * temp
		temp = temp * nums[i]
	}

	return result
}
