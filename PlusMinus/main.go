package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func plusMinus(arr []int32) {
	var p, n, z int32 = 0, 0, 0
	size := int32(len(arr))

	for _, num := range arr {
		if num > 0 { p++ }
		if num < 0 { n++ }
		if num == 0 { z++ }
	}

	fmt.Fprintf(os.Stdout, "%.6f\n", float32(p)/float32(size))
	fmt.Fprintf(os.Stdout, "%.6f\n", float32(n) / float32(size))
	fmt.Fprintf(os.Stdout, "%.6f\n", float32(z) / float32(size))
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

	plusMinus(arr)
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
