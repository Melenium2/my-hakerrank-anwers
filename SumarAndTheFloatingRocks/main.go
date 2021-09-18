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
	for a != b {
		if a > b {
			a, b = b, a
		}

		b = b - a
	}

	return a
}

func solve(x1 int32, y1 int32, x2 int32, y2 int32) int32 {
	if x1 == x2 {
		return int32(math.Abs(float64(y2-y1))) - 1
	} else if y1 == y2 {
		return int32(math.Abs(float64(x2-x1))) - 1
	} else {
		return gcd(int32(math.Abs(float64(y2-y1))), int32(math.Abs(float64(x2-x1)))) - 1
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		x1Y1X2Y2 := strings.Split(readLine(reader), " ")

		x1Temp, err := strconv.ParseInt(x1Y1X2Y2[0], 10, 64)
		checkError(err)
		x1 := int32(x1Temp)

		y1Temp, err := strconv.ParseInt(x1Y1X2Y2[1], 10, 64)
		checkError(err)
		y1 := int32(y1Temp)

		x2Temp, err := strconv.ParseInt(x1Y1X2Y2[2], 10, 64)
		checkError(err)
		x2 := int32(x2Temp)

		y2Temp, err := strconv.ParseInt(x1Y1X2Y2[3], 10, 64)
		checkError(err)
		y2 := int32(y2Temp)

		result := solve(x1, y1, x2, y2)

		fmt.Fprintf(writer, "%d\n", result)
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
