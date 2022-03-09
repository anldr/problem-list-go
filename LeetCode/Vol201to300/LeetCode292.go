package main

import "fmt"

func main() {
	fmt.Println(canWinNim(1))
	fmt.Println(canWinNim(2))
	fmt.Println(canWinNim(3))
	fmt.Println(canWinNim(4))
}

func canWinNim(n int) bool {
	if n%4 != 0 {
		return true
	} else {
		return false
	}
}
