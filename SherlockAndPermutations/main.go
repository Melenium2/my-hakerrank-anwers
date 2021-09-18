package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const mod = 1000000007

func factorial(n uint64) uint64 {
	if n > 0 {
		return (n * factorial(n-1)) % mod
	}
	return 1
}

func modpow(base, exp, modulus uint64) uint64 {
	base %= modulus
	var res uint64 = 1
	for exp > 0 {
		if exp & 1 == 1 {
			res = (res * base) % modulus
		}
		base = (base * base) % modulus
		exp >>= 1
	}

	return res
}

func solve(n int32, m int32) uint64 {
	div := (factorial(uint64(n)) % mod * factorial(uint64(m-1)) % mod) % mod
	div = modpow(div, mod-2, mod)

	return (factorial(uint64(n+m-1)) % mod * div % mod) % mod
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
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := solve(n, m)

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
