package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// a = b/2 * h
func findHeight(base, area int) float64 {
	return float64((2 * area - 1) / base + 1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	numbs := readLine(reader)
	split := strings.Split(numbs, " ")
	base, err := strconv.Atoi(split[0])
	checkError(err)
	area, err := strconv.Atoi(split[1])
	checkError(err)

	result := findHeight(base, area)
	fmt.Fprintf(os.Stdout, "%d", int64(result))
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
