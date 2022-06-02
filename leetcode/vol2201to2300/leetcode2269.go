package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(divisorSubstrings(240, 2))
	fmt.Println(divisorSubstrings(430043, 2))
}

func divisorSubstrings(num int, k int) int {
	var sub = ""
	var str = strconv.Itoa(num)

	var ret = 0
	for i := range str {
		if i >= k-1 {
			sub = sub + string(str[i])
			mod, _ := strconv.Atoi(sub)
			if mod > 0 && num%mod == 0 {
				ret++
			}
			sub = sub[1:]
		} else {
			sub = sub + string(str[i])
		}
	}

	return ret
}
