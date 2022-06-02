package main

import "fmt"
import . "../common_utils"

func main() {
	fmt.Println(largestCombination([]int{16, 17, 71, 62, 12, 24, 14}))
	fmt.Println(largestCombination([]int{8, 8}))
}

func largestCombination(candidates []int) int {
	var ret = 0
	for i := 0; i <= 30; i++ {
		tot := 0
		for j := range candidates {
			if candidates[j]&(1<<i) > 0 {
				tot++
			}
		}
		ret = MaxInt32(ret, tot)
	}

	return ret
}
