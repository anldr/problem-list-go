package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
	fmt.Println(countNumbersWithUniqueDigits(0))
}

func countNumbersWithUniqueDigits(n int) int {
	arr := []int{1, 9, 9, 8, 7, 6, 5, 4, 3}
	var result = 0

	var mutil = 1
	for i := 0; i <= n; i++ {
		mutil = mutil * arr[i]
		result += mutil
	}

	return result
}
