package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(1))
}

func numTrees(n int) int {
	catalan := make([]int, n+1)
	catalan[0] = 1
	catalan[1] = 1
	for i := 2; i <= n; i++ {
		var sum = 0
		for j := 1; j <= i; j++ {
			sum = sum + catalan[i-j]*catalan[j-1]
		}
		catalan[i] = sum
	}

	return catalan[n]
}
