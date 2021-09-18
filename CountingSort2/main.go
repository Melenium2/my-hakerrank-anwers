package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countingSort(arr []int32) []int32 {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] { max = arr[i] }
	}

	counting := make([]int32, max+1)
	for i := 0; i < len(arr); i++ {
		counting[arr[i]]++
	}

	for i := 1; i < len(counting); i++ {
		counting[i] += counting[i-1]
	}

	out := make([]int32, len(arr))
	for i := len(arr)-1; i >= 0; i-- {
		index := arr[i]
		out[counting[index]-1] = arr[i]
		counting[index]--
	}

	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024 * 15)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024 * 15)

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

	result := countingSort(arr)

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
