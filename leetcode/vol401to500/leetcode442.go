package main

import "fmt"

func main() {
	fmt.Println((findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})))
}

func findDuplicates(nums []int) (ans []int) {
	for i := 0; i < len(nums); i++ {
		if nums[Abs(nums[i])-1] < 0 {
			ans = append(ans, Abs(nums[i]))
		} else {
			nums[Abs(nums[i])-1] = -nums[Abs(nums[i])-1]
		}
	}

	return ans
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
