package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var pr = [16]int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

func primeCount(n int64) int {
	if n < 2 {
		return 0
	} else if n == 2 || n == 3 {
		return 1
	}
	var (
		it int
		sum int64 = 1
	)
	for i, v := range pr {
		sum *= int64(v)
		if sum <= n && i < 15{
			it++
		} else {
			break
		}
	}

	return it
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		result := primeCount(n)

		fmt.Fprintf(os.Stdout, "%d\n", result)
	}
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
