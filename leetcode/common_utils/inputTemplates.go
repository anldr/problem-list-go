package common_utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main1() {
	reader := bufio.NewReader(os.Stdin)
	ints, err := readIntSlice(reader, " ")
	if err != nil {
		return
	}
	threshold := ints[0]
	usersPerHour, err := readIntSlice(reader, " ")
	if err != nil {
		return
	}
	fmt.Print(usersPerHour, threshold)
}

func readIntSlice(reader *bufio.Reader, sep string) ([]int, error) {
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf(err.Error())
	}

	lineBuf = strings.TrimRight(lineBuf, "\r\n")
	line := strings.Split(lineBuf, sep)
	var result []int
	for _, v := range line {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		result = append(result, int(i))
	}
	return result, nil
}

func main2() {
	inputReader := bufio.NewReader(os.Stdin)
	diff, inputs, err := readInput1793(inputReader)
	if err != nil {
		return
	}
	fmt.Println(inputs, diff)
}

func readInput1793(reader *bufio.Reader) (int, []int, error) {
	var s int
	if _, err := fmt.Fscanf(reader, "%d\n", &s); err != nil {
		return 0, nil, err
	}

	var num int
	if _, err := fmt.Fscanf(reader, "%d\n", &num); err != nil {
		return 0, nil, err
	}
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, nil, err
	}
	lineBuf = strings.TrimRight(lineBuf, "\r\n")
	inputsArray := strings.Split(lineBuf, " ")
	if num != len(inputsArray) {
		return 0, nil, fmt.Errorf("%s", "input error")
	}
	inputs := make([]int, num, num)
	for index, val := range inputsArray {
		intVal, _ := strconv.Atoi(val)
		inputs[index] = int(intVal)
	}

	return s, inputs, nil
}
