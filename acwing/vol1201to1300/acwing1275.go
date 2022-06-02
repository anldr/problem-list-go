package main

import (
	. "../../leetcode/common_utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SegeMentTree struct {
	maxx         int64
	lRng, rRng   int
	lNode, rNode *SegeMentTree
}

func (tr *SegeMentTree) update() {
}

func (tr *SegeMentTree) pushUp() {
	tr.maxx = MaxInt64(tr.maxx, MaxInt64(tr.lNode.maxx, tr.rNode.maxx))
}

func (tr *SegeMentTree) pushDown() {
}

func (tr *SegeMentTree) build(l, r int) {
	tr.lRng, tr.rRng = l, r
	if l == r {
		tr.maxx = int64(-1)
	} else {
		mid := (l + r) >> 1
		tr.lNode = &SegeMentTree{}
		tr.lNode.build(l, mid)
		tr.rNode = &SegeMentTree{}
		tr.rNode.build(mid+1, r)
		tr.pushUp()
	}
}

func (tr *SegeMentTree) modify(l, r int, val int64) {
	L, R := tr.lRng, tr.rRng
	if l <= L && R <= r {
		tr.maxx = val
	} else {
		tr.pushDown()
		mid := (L + R) >> 1
		if mid >= l {
			tr.lNode.modify(l, r, val)
		}
		if r > mid {
			tr.rNode.modify(l, r, val)
		}
		tr.pushUp()
	}
}

func (tr *SegeMentTree) query(l, r int) int64 {
	L, R := tr.lRng, tr.rRng
	if l <= L && R <= r {
		return tr.maxx
	} else {
		tr.pushDown()
		var ans = int64(0)
		mid := (L + R) >> 1
		if mid >= l {
			ans = tr.lNode.query(l, r)
		}
		if r > mid {
			ans = MaxInt64(ans, tr.rNode.query(l, r))
		}
		tr.pushUp()
		return ans
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	m, p := readInteger(reader, " ")

	var add = int64(0)
	var totLen = 0
	tr := SegeMentTree{}
	tr.build(1, 200007)
	for i := 0; i < m; i++ {
		op, num := readLine(reader, " ")
		if op == 'A' {
			tr.modify(totLen+1, totLen+1, (add+num)%p)
			totLen++
		} else {
			add = tr.query(totLen-int(num)+1, totLen)
			fmt.Println(add)
		}
	}
}

func readInteger(reader *bufio.Reader, sep string) (int, int64) {
	lineBuf, _ := reader.ReadString('\n')
	lineBuf = strings.Trim(lineBuf, "\r\n")
	line := strings.Split(lineBuf, sep)

	m, _ := strconv.ParseInt(line[0], 10, 32)
	p, _ := strconv.ParseInt(line[1], 10, 32)

	return int(m), p
}

func readLine(reader *bufio.Reader, sep string) (byte, int64) {
	lineBuf, _ := readInput(reader)
	lineBuf = strings.Trim(lineBuf, "\r\n")
	line := strings.Split(lineBuf, sep)

	ch := line[0][0]
	num, _ := strconv.ParseInt(line[1], 10, 32)

	return ch, num
}

func readInput(reader *bufio.Reader) (string, error) {
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return lineBuf, nil
}
