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

func partition(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}

	pivot := arr[0]
	l := make([]int32, 0)
	r := make([]int32, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}

	if len(l) > 1 {
		l = partition(l)
	}
	if len(r) > 1 {
		r = partition(r)
	}

	l = append(l, pivot)
	l = append(l, r...)
	printArray(l)

	return l
}

func quickSort(arr []int32) {
	partition(arr)
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

	quickSort(arr)
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
