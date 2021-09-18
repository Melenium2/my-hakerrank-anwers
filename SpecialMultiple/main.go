package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(n int64) string {
	var nines []string
	nines = append(nines, "9")
	i := 0
	sl := 0
	for {
		first := nines[i]
		second := first

		first += "0"
		nines = append(nines, first)
		second += "9"
		nines = append(nines, second)
		for _, v := range nines[sl:sl+3] {
			num, _ := strconv.ParseInt(v, 10, 64)
			if num % n == 0 {
				return v
			}
		}
		i++
		sl += 2
	}
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

		result := solve(nTemp)

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
