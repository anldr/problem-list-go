package main

import (
	. "../common_utils"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumWhiteTiles([][]int{{10, 11}, {1, 1}}, 2))
}

type tileRng struct {
	left  int
	right int
}

type tileRngs []tileRng

func (ts tileRngs) Len() int {
	return len(ts)
}

func (ts tileRngs) Less(i, j int) bool {
	return ts[i].left < ts[j].left
}

func (ts tileRngs) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	var arrLen = len(tiles)
	var tsSort = make(tileRngs, arrLen)
	for i := range tiles {
		tsSort[i].left = tiles[i][0]
		tsSort[i].right = tiles[i][1]
	}
	sort.Sort(tsSort)

	var maxLen = -1
	for i := 0; i < arrLen; i++ {
		var temp = 0
		var left = tsSort[i].left
		var k = i
		for ; k < arrLen && left+carpetLen-1 >= tsSort[k].right; k++ {
			temp = temp + tsSort[k].right - tsSort[k].left + 1
		}
		if k >= arrLen {
			maxLen = MaxInt32(maxLen, temp)
			break
		} else {
			if tsSort[k].left <= left+carpetLen-1 {
				temp = temp + left + carpetLen - tsSort[k].left
			}
			maxLen = MaxInt32(maxLen, temp)
		}

		if maxLen >= carpetLen {
			break
		}
	}

	return MinInt32(maxLen, carpetLen)
}
