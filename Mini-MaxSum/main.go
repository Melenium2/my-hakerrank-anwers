package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int64) {
	sums := make([]int64, 0)
	var max, sum int64 = 0, 0
	for i, _ := range arr {
		for j := 0; j < len(arr); j++ {
			if j == i { continue }
			sum += arr[j]
		}
		sums = append(sums, sum)
		if max == 0 || max < sum {
			max = sum
		}
		sum = 0
	}

	var min int64 = sums[0]
	for i := 1; i < len(sums); i++ {
		if min > sums[i] {
			min = sums[i]
		}
	}

	fmt.Fprintf(os.Stdout, "%d %d\n", min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
