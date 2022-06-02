package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
	fmt.Println(largestGoodInteger("2300019"))
	fmt.Println(largestGoodInteger("42352338"))
	fmt.Println(largestGoodInteger("222"))
}

func largestGoodInteger(num string) string {
	var ret = ""

	pre := num[0]
	var prestr strings.Builder
	for i := range num {
		if num[i] == pre {
			prestr.WriteString(string(num[i]))
		} else {
			if len(prestr.String()) >= 3 {
				substr := prestr.String()[len(prestr.String())-3 : len(prestr.String())]
				if ret < substr {
					ret = substr
				}
			}
			prestr.Reset()
			prestr.WriteString(string(num[i]))
		}
		pre = num[i]
	}

	if len(prestr.String()) >= 3 {
		substr := prestr.String()[len(prestr.String())-3 : len(prestr.String())]
		if ret < substr {
			ret = substr
		}
	}

	return ret
}
