package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))
}

func uniqueMorseRepresentations(words []string) int {
	key := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	var result = 0
	set := make(map[string]int)
	for _, str := range words {
		strBuilder := strings.Builder{}
		for _, ch := range str {
			strBuilder.WriteString(key[ch-'a'])
		}

		if _, ok := set[strBuilder.String()]; !ok {
			result++
			set[strBuilder.String()]++
		}
	}

	return result
}
