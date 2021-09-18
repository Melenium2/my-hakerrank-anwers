package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printArray(arr []int32) {
	for _, n := range arr {
		fmt.Fprintf(os.Stdout, "%d ", n)
	}
	fmt.Fprint(os.Stdout, "\n")
}

func partition(arr []int32, l, r int) int {
	pivot := arr[r]
	i := l

	for j := l; j <= r; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	printArray(arr)

	return i
}

func quickSort(arr []int32, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, p - 1)
		quickSort(arr, p + 1, r)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

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

	quickSort(arr, 0, len(arr) - 1)
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
