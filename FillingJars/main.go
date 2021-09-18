package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(n int32, operations [][]int64) int64 {
	var sum int64 = 0
	for _, v := range operations {
		sum += (v[1] - v[0] + 1) * v[2]
	}
	return sum/int64(n)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var operations [][]int64
	for operationsRowItr := 0; operationsRowItr < int(m); operationsRowItr++ {
		operationsRowTemp := strings.Split(readLine(reader), " ")

		var operationsRow []int64
		for _, operationsRowItem := range operationsRowTemp {
			operationsItemTemp, err := strconv.ParseInt(operationsRowItem, 10, 64)
			checkError(err)
			operationsItem := operationsItemTemp
			operationsRow = append(operationsRow, operationsItem)
		}

		if len(operationsRow) != int(3) {
			panic("Bad input")
		}

		operations = append(operations, operationsRow)
	}

	result := solve(n, operations)

	fmt.Fprintf(writer, "%d\n", result)

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
