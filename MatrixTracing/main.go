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

var mod = int64(math.Pow(10, 9) + 7)

func f(n int64) int64 {
	if n > 0 {
		return (n * f(n-1))%mod
	}

	return 1
}

func modpow(b, e, m int64) int64 {
	b %= m
	var res int64 = 1
	for e > 0 {
		if e & 1 == 1 {
			res = (res * b) % m
		}
		b = (b*b) % m
		e >>= 1
	}

	return res
}

func solve(n int64, m int64) int64 {
	num := f(m+n - 2)
	denom := f(n-1) * f(m-1)
	return ((num%mod) * modpow(denom, mod-2, mod))%mod
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)

		result := solve(nTemp, mTemp)

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
