package main

import (
	. "../common_struct"
	"fmt"
)

func main() {
	fmt.Println(countTexts("22233"))
	fmt.Println(countTexts("222222222222222222222222222222222222"))
}

var dict = [][]byte{{'x'}, {'x'}, {'a', 'b', 'c'}, {'d', 'e', 'f'},
	{'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'},
	{'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
var mod = 1000000007

func countTexts(pressedKeys string) int {
	var strLen = len(pressedKeys)
	var dp = make([]int, strLen+1)

	var tot = 0
	var pre = pressedKeys[0]
	dp[0] = 1
	for i := 1; i <= strLen; i++ {
		if pre == pressedKeys[i-1] {
			tot++
			var arrLen = len(dict[pre-'0'])
			for k := 1; k <= Min(arrLen, tot); k++ {
				dp[i] = (dp[i] + dp[i-k]) % mod
			}
		} else {
			tot = 1
			dp[i] = dp[i-1] % mod
		}
		pre = pressedKeys[i-1]
	}

	return dp[strLen]
}
