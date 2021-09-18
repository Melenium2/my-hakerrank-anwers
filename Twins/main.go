package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int32) bool {
	if n <= 1 { return false }
	if n <= 3 { return true }

	if n % 2 == 0 || n % 3 == 0 { return false }

	var i int32 = 0
	for i = 5; i*i <= n; i = i + 6 {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
	}

	return true
}

func solve(n int32, m int32) int32 {
	var counter int32 = 0
	for i := n; i <= m-2; i+=2 {
		if isPrime(i) && isPrime(i+2) && i+2 - i == 2 {
			counter++
		}
	}

	return counter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	if n%2 != 1 {
		n += 1
	}
	result := solve(n, m)

	fmt.Fprintf(writer, "%d\n", result)

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
