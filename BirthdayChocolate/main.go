package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sum(s []int32, i, m int32) int32 {
	var sum int32 = 0
	for j := i; j < i+m; j++ {
		sum += s[j]
	}
	return sum
}

func birthday(s []int32, d int32, m int32) int32 {
	var counter int32 = 0
	for i := 0; i <= len(s) - int(m); i++ {
		sum := sum(s, int32(i), m)
		if sum == d { counter++ }
	}

	return counter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	dm := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	dTemp, err := strconv.ParseInt(dm[0], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(dm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := birthday(s, d, m)

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
