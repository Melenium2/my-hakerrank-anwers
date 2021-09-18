package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func f(x, y uint64, a []uint64) string {
	if a[x-1]%2 == 0 && (x==y || a[x]!= 0) {
		return "Even"
	}
	return "Odd"
}

func solve(a []uint64, queries [][]uint64) []string {
	var res []string
	for _, q := range queries {
		r := f(q[0], q[1], a)
		res = append(res, r)
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	arrCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arrTemp := strings.Split(readLine(reader), " ")

	var a []uint64

	for arrItr := 0; arrItr < int(arrCount); arrItr++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
		checkError(err)
		a = append(a, uint64(arrItemTemp))
	}

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]uint64
	for queriesRowItr := 0; queriesRowItr < int(q); queriesRowItr++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []uint64
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := queriesItemTemp
			queriesRow = append(queriesRow, uint64(queriesItem))
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := solve(a, queries)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if resultItr != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
