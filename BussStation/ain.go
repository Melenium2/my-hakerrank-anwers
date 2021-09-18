package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(a []int) []int {
	var sum, i int
	for _, n := range a {
		sum += n
	}
	var divs []int
	for i = 1; i*i <= sum; i++ {
		if sum%i == 0 {
			divs = append(divs, i)
			if i*i != sum {
				divs = append(divs, sum/i)
			}
		}
	}

	sort.Ints(divs)

	var res []int
	for _, d := range divs {
		var subsum = 0
		i = 0
		for subsum <= d && i < len(a) {
			subsum += a[i]
			i++
			if subsum == d {
				subsum = 0
			}
		}
		if subsum <= d {
			res = append(res, d)
		}
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	aCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int

	for aItr := 0; aItr < int(aCount); aItr++ {
		aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
		checkError(err)
		aItem := int(aItemTemp)
		a = append(a, aItem)
	}

	result := solve(a)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result) - 1 {
			fmt.Fprintf(writer, " ")
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
