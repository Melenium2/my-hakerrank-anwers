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

func solve(n int64, m int64) int64 {
	return big.NewInt(n).Mul(big.NewInt(n), big.NewInt(m)).Int64() - 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	nm := strings.Split(readLine(reader), " ")
	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int64(nTemp)
	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int64(mTemp)

	result := solve(n, m)

	fmt.Fprintf(os.Stdout, "%d\n", result)
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
