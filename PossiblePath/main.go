package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int64) int64 {
	for b > 0 {
		a %= b
		a, b = b, a
	}

	return a
}

func solve(a int64, b int64, x int64, y int64) string {
	if gcd(a, b) == gcd(x ,y) {
		return "YES"
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		abxy := strings.Split(readLine(reader), " ")

		a, err := strconv.ParseInt(abxy[0], 10, 64)
		checkError(err)

		b, err := strconv.ParseInt(abxy[1], 10, 64)
		checkError(err)

		x, err := strconv.ParseInt(abxy[2], 10, 64)
		checkError(err)

		y, err := strconv.ParseInt(abxy[3], 10, 64)
		checkError(err)

		result := solve(a, b, x, y)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
