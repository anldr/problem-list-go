package main

import "fmt"

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	fmt.Println(calcEquation(equations, values, queries))
}

var unionP []int
var unionV []float64

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var idx = 1
	var mapNum = map[string]int{}

	for _, v := range equations {
		for i := 0; i <= 1; i++ {
			if _, ok := mapNum[v[i]]; !ok {
				mapNum[v[i]] = idx
				idx++
			}
		}
	}

	var nodeNum = len(mapNum)
	unionP = make([]int, nodeNum+1)
	unionV = make([]float64, nodeNum+1)
	for i := 1; i <= nodeNum; i++ {
		unionP[i] = i
		unionV[i] = 1.0
	}

	for i, v := range equations {
		union(mapNum[v[0]], mapNum[v[1]], values[i])
	}

	var result = []float64{}
	for _, v := range queries {
		num1, ok1 := mapNum[v[0]]
		num2, ok2 := mapNum[v[1]]
		if !ok1 || !ok2 {
			result = append(result, -1.00000)
			continue
		}
		if num1 == num2 {
			result = append(result, 1.00000)
			continue
		}
		result = append(result, getVal(num1, num2))
	}

	return result
}

func union(x int, y int, val float64) {
	var root1 = find(x)
	var root2 = find(y)

	if root1 != root2 {
		unionP[root1] = root2
		unionV[root1] = unionV[y] * val / unionV[x]
	}
}

func find(x int) int {
	if unionP[x] != x {
		var temp = find(unionP[x])
		unionV[x] = unionV[x] * unionV[unionP[x]]
		unionP[x] = temp
	}

	return unionP[x]
}

func getVal(x int, y int) float64 {
	var root1 = find(x)
	var root2 = find(y)
	if root1 == root2 {
		return unionV[x] / unionV[y]
	}

	return -1.00000
}
