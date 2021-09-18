package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func modpow(base, exp, modulus int64) int64 {
	base %= modulus
	var res int64 = 1
	for exp > 0 {
		if exp&1 == 1 {
			res = (res*base) % modulus
		}
		exp >>= 1
		base = (base*base) % modulus
	}

	return res
}

func solve(a int64, m int64) string {
	if a == 0 { return "YES" }
	if modpow(a, (m-1)/2, m) == 1 { return "YES" }

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := tTemp

	for tItr := 0; tItr < int(t); tItr++ {
		am := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(am[0], 10, 64)
		checkError(err)
		a := aTemp

		mTemp, err := strconv.ParseInt(am[1], 10, 64)
		checkError(err)
		m := mTemp

		result := solve(a, m)

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
