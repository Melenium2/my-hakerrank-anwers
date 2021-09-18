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

var mod = uint64(math.Pow(10, 9))

func triangle(n, k uint64) uint64 {
	t := make([][]uint64, n+1)
	for in, _ := range t {
		t[in] = make([]uint64, k+1)
	}
	var i, j uint64

	for i = 0; i <= n; i++ {
		for j = 0; j <= uint64(math.Min(float64(i), float64(k))); j++ {
			if i == 0 || j == 0 {
				t[i][j] = 1
			} else {
				t[i][j] = (t[i-1][j-1] + t[i-1][j]) % mod
			}
		}
	}

	return t[n][k]%mod
}

func solve(n uint64, k uint64) uint64 {
	return triangle(n+k-1, k)
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
		n := uint64(nTemp)

		kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		k := uint64(kTemp)

		result := solve(n, k)

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
