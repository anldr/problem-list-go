package main

import (
	. "../common_utils"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxConsecutive(2, 9, []int{4, 6}))
	fmt.Println(maxConsecutive(3, 15, []int{7, 9, 13}))
}

func maxConsecutive(bottom int, top int, special []int) int {
	var ret = 0

	special = append(special, bottom)
	special = append(special, top)
	var arrLen = len(special)
	sort.Ints(special)

	for i := 1; i < arrLen; i++ {
		if i == 1 || i == arrLen-1 {
			ret = MaxInt32(ret, special[i]-special[i-1])
		} else {
			ret = MaxInt32(ret, special[i]-special[i-1]-1)
		}
	}

	return ret
}
