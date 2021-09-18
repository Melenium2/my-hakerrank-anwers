package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a % b
	}

	return a
}

func solve(a []int32) string {
	g := a[0]
	for i := 1; i < len(a); i++ {
		g = gcd(g, a[i])
	}

	if g > 1 { return "NO" } else { return "YES" }
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		aCount, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		aTemp := strings.Split(readLine(reader), " ")

		var a []int32

		for aItr := 0; aItr < int(aCount); aItr++ {
			aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			a = append(a, aItem)
		}

		result := solve(a)

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
