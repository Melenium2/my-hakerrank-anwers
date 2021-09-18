package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(dates [][]int32) int {
	events := make(map[int]int, 0)
	for _, dateRow := range dates {
		m := int(dateRow[0])
		d := int(dateRow[1])
		dateString := strconv.FormatInt(int64(d), 10)
		decimalEvent, err := strconv.ParseInt(dateString, m, 64)
		if err == nil {
			events[int(decimalEvent)]++
		}
	}
	output := 0
	for _, v := range events {
		output += v * (v - 1) / 2
	}
	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var dates [][]int32
	for datesRowItr := 0; datesRowItr < int(n); datesRowItr++ {
		datesRowTemp := strings.Split(readLine(reader), " ")

		var datesRow []int32
		for _, datesRowItem := range datesRowTemp {
			datesItemTemp, err := strconv.ParseInt(datesRowItem, 10, 64)
			checkError(err)
			datesItem := int32(datesItemTemp)
			datesRow = append(datesRow, datesItem)
		}

		if len(datesRow) != int(2) {
			panic("Bad input")
		}

		dates = append(dates, datesRow)
	}

	result := solve(dates)

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