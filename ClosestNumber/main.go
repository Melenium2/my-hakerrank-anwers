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

func closestNumber(a int32, b int32, x int32) int32 {
	expect := math.Pow(float64(a), float64(b))
	xf := float64(x)

	return int32(math.Round(expect/xf) * xf)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		abx := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(abx[0], 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(abx[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		xTemp, err := strconv.ParseInt(abx[2], 10, 64)
		checkError(err)
		x := int32(xTemp)

		result := closestNumber(a, b, x)

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
