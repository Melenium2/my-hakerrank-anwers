package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int32) int32 {
	for a%b != 0 {
		a, b = b, a%b
	}

	return b
}

func solve(a int32, b int32, c int32) string {
	if float64(c) <= math.Max(float64(a),float64(b)) && c%gcd(a,b) == 0 {
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
		abc := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(abc[0], 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(abc[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		cTemp, err := strconv.ParseInt(abc[2], 10, 64)
		checkError(err)
		c := int32(cTemp)

		result := solve(a, b, c)

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
