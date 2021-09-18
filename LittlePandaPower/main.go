package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mmi(a, m int64) int64 {
	var m0 = m
	var x, y int64 = 1, 0
	if m == 1 {
		return 1
	}
	for a > 1 {
		q := a / m
		a, m = m, a % m
		x, y = y, x - q * y
	}

	if x < 0 {
		x += m0
	}

	return x
}

func modpow(base, exp, mod int64) int64 {
	base = base % mod
	var res int64 = 1
	for exp > 0 {
		if exp & 1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}

	return res
}

func solve(a, b, x int64) int64 {
	if b > 0 {
		return modpow(a, b, x)
	}

	res := mmi(a, x)

	return modpow(res, b*-1, x)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		abx := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(abx[0], 10, 64)
		checkError(err)
		a := aTemp

		bTemp, err := strconv.ParseInt(abx[1], 10, 64)
		checkError(err)
		b := bTemp

		xTemp, err := strconv.ParseInt(abx[2], 10, 64)
		checkError(err)
		x := xTemp

		result := solve(a, b, x)

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
