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

func insertionSort2(n int32, arr []int32) {
	for i, j := 0, 1; j < int(n); i, j = i + 1, j + 1 {
		if arr[j] < arr[i] {
			el := arr[j]
			k := i
			for k >=0 && el < arr[k] {
				arr[k+1] = arr[k]
				k--
			}
			arr[k+1] = el
		}
		printArray(arr)
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

	insertionSort2(n, arr)
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
