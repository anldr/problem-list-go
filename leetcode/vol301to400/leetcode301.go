package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")("))
}

func removeInvalidParentheses(s string) []string {
	var leftPr []int
	var rightPr []int
	leftPar, rightPar := 0, 0
	for i, v := range s {
		if v == '(' {
			leftPr = append(leftPr, i)
			leftPar++
		} else if v == ')' {
			rightPr = append(rightPr, i)
			if leftPar == 0 {
				rightPar++
			} else {
				leftPar--
			}
		}
	}

	var leftLen = len(leftPr)
	var rightLen = len(rightPr)
	var leftMaskArr []int
	var rightMaskArr []int
	for i := 0; i < (1 << leftLen); i++ {
		if bits.OnesCount(uint(i)) == leftPar {
			leftMaskArr = append(leftMaskArr, i)
		}
	}
	for i := 0; i < (1 << rightLen); i++ {
		if bits.OnesCount(uint(i)) == rightPar {
			rightMaskArr = append(rightMaskArr, i)
		}
	}

	result := map[string]struct{}{}
	for _, leftMask := range leftMaskArr {
		for _, rightMask := range rightMaskArr {
			ok, str := isValid(s, leftMask, rightMask, leftPr, rightPr)
			if ok {
				result[str] = struct{}{}
			}
		}
	}

	ret := []string{}
	for k, _ := range result {
		ret = append(ret, k)
	}

	return ret
}

func isValid(str string, leftMask, rightMask int, leftPr, rightPr []int) (bool, string) {
	count := 0
	ret := []rune{}
	leftPos, rightPos := 0, 0
	var leftLen = len(leftPr)
	var rightLen = len(rightPr)
	for i, ch := range str {
		if leftPos < leftLen && i == leftPr[leftPos] {
			if (leftMask >> leftPos & 1) == 0 {
				count++
				ret = append(ret, '(')
			}
			leftPos++
		} else if rightPos < rightLen && i == rightPr[rightPos] {
			if (rightMask >> rightPos & 1) == 0 {
				count--
				ret = append(ret, ')')
				if count < 0 {
					return false, "xxx"
				}
			}
			rightPos++
		} else {
			ret = append(ret, ch)
		}
	}

	return count == 0, string(ret)
}

func recoverString(str string, leftMask, rightMask int, leftPr, rightPr []int) string {
	ret := []rune{}
	leftPos, rightPos := 0, 0
	var leftLen = len(leftPr)
	var rightLen = len(rightPr)
	for i, ch := range str {
		if leftPos < leftLen && i == leftPr[leftPos] {
			if (leftMask >> leftPos & 1) == 0 {
				ret = append(ret, '(')
			}
			leftPos++
		} else if rightPos < rightLen && i == rightPr[rightPos] {
			if (rightMask >> rightPos & 1) == 0 {
				ret = append(ret, ')')
			}
			rightPos++
		} else {
			ret = append(ret, ch)
		}
	}

	return string(ret)
}
