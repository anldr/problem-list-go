package main

import "fmt"

func main() {
	fmt.Println(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}))
	fmt.Println(removeAnagrams([]string{"a", "b", "c", "d", "e"}))
}

func removeAnagrams(words []string) (ans []string) {
	var arrLen = len(words)
	var left = 0
	var right = 1

	var visit = make([]bool, arrLen)
	for left < arrLen && right < arrLen {
		if visit[left] == true {
			continue
		}
		for right < arrLen {
			if words[left] == words[right] || isEq(words[left], words[right]) {
				visit[right] = true
				right++
			} else {
				left = right
				right++
				break
			}
		}
	}

	ans = []string{}
	for i := range visit {
		if visit[i] == false {
			ans = append(ans, words[i])
		}
	}

	return ans
}

func isEq(str1 string, str2 string) bool {
	var arr1 = make([]int, 26)
	var arr2 = make([]int, 26)
	for _, v := range str1 {
		arr1[v-'a'] = arr1[v-'a'] + 1
	}
	for _, v := range str2 {
		arr2[v-'a'] = arr2[v-'a'] + 1
	}

	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
