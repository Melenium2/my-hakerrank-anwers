package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the canConstruct function below.
func canConstruct(a []int32, division int32) string {
	var numbs []int32
	for _, v := range a {
		for v > 0 {
			digit := v % 10
			v = (v - digit) / 10
			numbs = append(numbs, digit)
		}
	}

	var sum int32
	for _, v := range numbs {
		sum += v
	}
	if sum % 3 == 0 {
		return "Yes"
	}
	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	division := int32(tTemp)

	for tItr := 0; tItr < int(division); tItr++ {
		count, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(count)

		numbs := strings.Split(readLine(reader), " ")

		var a []int32

		for i := 0; i < int(n); i++ {
			aItemTemp, err := strconv.ParseInt(numbs[i], 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			a = append(a, aItem)
		}

		result := canConstruct(a, division)

		fmt.Fprintf(os.Stdout, "%s\n", result)
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
