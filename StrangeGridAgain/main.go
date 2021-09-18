package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func strangeGrid(r int32, c int32) int64 {
	var (
		ten = int64(r/2)
		digit int32 = 0
		even = []int32{0, 2, 4, 6, 8}
		odd = []int32{1, 3, 5, 7, 9}
	)
	if r % 2 == 0 {
		ten -= 1
		digit = odd[c - 1]
	} else {
		digit = even[c - 1]
	}
	return ten*10 + int64(digit)
}

func anotherSolution(r int32, c int32) int64 {
	return int64(((r-1)/2) * 10 + 2 * (c-1) + (r-1)%2)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	rc := strings.Split(readLine(reader), " ")

	rTemp, err := strconv.ParseInt(rc[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	cTemp, err := strconv.ParseInt(rc[1], 10, 64)
	checkError(err)
	c := int32(cTemp)

	result := anotherSolution(r, c)

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
