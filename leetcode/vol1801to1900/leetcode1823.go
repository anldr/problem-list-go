package main

import "fmt"

func main() {
	fmt.Println(findTheWinner(5, 2))
	fmt.Println(findTheWinner(6, 5))
}

/*
约瑟夫环推理：
	1.
*/
func findTheWinner(n int, k int) int {
	var result = 1
	for i := 2; i <= n; i++ {
		result = (k+result-1)%i + 1
	}

	return result
}
