package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	var sb strings.Builder
	for _, v := range s {
		if v == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteByte(byte(v))
		}
	}

	return sb.String()
}
