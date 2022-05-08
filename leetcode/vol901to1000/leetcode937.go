package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
}

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s, t := logs[i], logs[j]

		s0 := strings.SplitN(s, " ", 2)[1]
		t0 := strings.SplitN(t, " ", 2)[1]
		sDig := unicode.IsDigit(rune(s0[0]))
		tDig := unicode.IsDigit(rune(t0[0]))

		if sDig && tDig {
			return false
		}
		if !sDig && !tDig {
			return s0 < t0 || s0 == t0 && s < t
		}

		return !sDig
	})

	return logs
}
