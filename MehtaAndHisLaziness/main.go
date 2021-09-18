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
func perfectSquare(n int64) int64 {
	if n <= 2 || n % 2 != 0 { return 0 }
	sq := math.Sqrt(float64(n))
	if (sq - math.Floor(sq)) == 0 {
		return 1
	}

	return 0
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a % b
	}

	return a
}

func solve(n int64) string {
	var i, d, p int64 = 0, 1, 0
	for i = 1; i <= int64(math.Sqrt(float64(n))); i++ {
		if n % i == 0 && n/i != n {
			d++; p+=perfectSquare(i)
			if n/i == i {
			} else {
				d++; p+=perfectSquare(n/i)
			}
		}
	}
	if p == 0 { return "0" }

	k := gcd(p, d)
	return fmt.Sprintf("%d/%d", p/k, d/k)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := nTemp

		result := solve(n)

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
