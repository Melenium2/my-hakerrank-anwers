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

var mod = int64(math.Pow10(9) + 7)

func intiFibMatrix(a, b int64) [2][2]int64 {
	var matrix [2][2]int64
	matrix[0][0] = b
	matrix[0][1] = a
	matrix[1][0] = a
	matrix[1][1] = 0

	return matrix
}

func pow(f *[2][2]int64, n int64) {
	if n == 0 || n == 1 {
		return
	}
	m := intiFibMatrix(1, 1)

	pow(f, n / 2)
	multiply(f, f)

	if n%2 != 0 {
		multiply(f, &m)
	}
}

func multiply(f, m *[2][2]int64){
	x := (f[0][0] * m[0][0]) % mod + (f[0][1] * m[1][0]) % mod
	y := (f[0][0] * m[0][1]) % mod + (f[0][1] * m[1][1]) % mod
	z := (f[1][0] * m[0][0]) % mod + (f[1][1] * m[1][0]) % mod
	w := (f[1][0] * m[0][1]) % mod + (f[1][1] * m[1][1]) % mod

	f[0][0] = x
	f[0][1] = y
	f[1][0] = z
	f[1][1] = w
}

func solve(a , b , n int64) int64 {
	m := intiFibMatrix(1, 1)

	pow(&m, n - 1)

	return ((m[0][0] * b%mod) + (m[0][1] * a%mod)) % mod
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		abn := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(abn[0], 10, 64)
		checkError(err)
		a := aTemp

		bTemp, err := strconv.ParseInt(abn[1], 10, 64)
		checkError(err)
		b := bTemp

		nTemp, err := strconv.ParseInt(abn[2], 10, 64)
		checkError(err)
		n := nTemp

		result := solve(a, b, n)

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
