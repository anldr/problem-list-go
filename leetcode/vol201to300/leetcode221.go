package main

import "math"

func main() {

}

func maximalSquare(matrix [][]byte) int {
	var rowLen = len(matrix)
	var colLen = len(matrix[0])

	row := make([]int, colLen)

	var result = 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if matrix[i][j] == '0' {
				row[j] = 0
			} else {
				row[j] = row[j] + int(matrix[i][j]-'0')
			}
		}
		result = Max(result, largestArea(row))
	}

	return result
}

func largestArea(heights []int) int {
	newHeight := []int{}
	newHeight = append(newHeight, -1)
	newHeight = append(newHeight, heights...)
	newHeight = append(newHeight, -1)

	arrLen := len(newHeight)
	arrLen2 := len(heights)

	stack := []int{}
	leftMin := make([]int, arrLen2)
	stack = append(stack, 0)
	for i := 1; i < arrLen-1; i++ {
		for newHeight[stack[len(stack)-1]] >= newHeight[i] {
			stack = stack[:len(stack)-1]
		}
		leftMin[i-1] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	stack = []int{}

	rightMin := make([]int, arrLen2)
	stack = append(stack, arrLen-1)
	for i := arrLen - 2; i >= 1; i-- {
		for newHeight[stack[len(stack)-1]] >= newHeight[i] {
			stack = stack[:len(stack)-1]
		}
		rightMin[i-1] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	var result = 0
	for i := 0; i < arrLen2; i++ {
		var temp = Min(rightMin[i]-leftMin[i]-1, heights[i])
		result = int(math.Max(float64(result), float64(temp*temp)))
	}

	return result
}

func Max(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}

func Min(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}
