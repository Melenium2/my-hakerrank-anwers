package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func toString(el, n int) string {
	str := ""
	for i := 0; i < el; i++ {
		str += "#"
	}
	for i := len(str); i < n; i++ {
		str = " " + str
	}

	return str
}

func staircase(n int32) {
	s := int(n)
	for i := 1; i <= s; i++ {
		str := toString(i, s)
		fmt.Fprintf(os.Stdout, "%s\n", str)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
