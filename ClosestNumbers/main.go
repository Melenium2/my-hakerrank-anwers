package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int32, l, r int) int {
	pivot := arr[r]
	i := l

	for j := l; j < r; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[i], arr[r] = arr[r], arr[i]

	return i
}

func abs(a int32) int32 {
	return int32(math.Abs(float64(a)))
}

func quickSort(arr []int32, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
}

func closestNumbers(arr []int32) []int32 {
	if len(arr) < 3 {
		return []int32{abs(arr[0] - arr[1])}
	}

	quickSort(arr, 0, len(arr)-1)

	min := abs(arr[0] - arr[1])
	mins := make([]int32, 0)
	mins = append(mins, arr[0], arr[1])
	for i := 1; i < len(arr)-1; i++ {
		v := abs(arr[i] - arr[i+1])
		if min > v {
			min = v
			mins = nil
			mins = append(mins, arr[i], arr[i+1])
		} else if min == v {
			mins = append(mins, arr[i], arr[i+1])
		}
	}

	return mins
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := closestNumbers(arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
