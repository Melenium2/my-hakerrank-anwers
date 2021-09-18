package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func gameWithCells(n int32, m int32) int32 {
	if n % 2 != 0 {
		n++
	}
	if m % 2 != 0 {
		m++
	}
	return n*m / 4
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := gameWithCells(n, m)

	fmt.Fprintf(os.Stdout, "%d\n", result)
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
