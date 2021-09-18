package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func summingSeries(n int64) int64 {
	return big.NewInt(n).Mod(big.NewInt(n).Mul(big.NewInt(n), big.NewInt(n)), big.NewInt(1000000007)).Int64()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		result := summingSeries(n)

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
