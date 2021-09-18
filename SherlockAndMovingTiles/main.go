package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func movingTiles(l int32, s1 int32, s2 int32, queries []int64) []float64 {
	result := make([]float64, 0)
	for _, v := range queries {
		result = append(result, math.Sqrt(2) * (float64(l) - math.Sqrt(float64(v))) / math.Abs(float64(s2 - s1)) )
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	lS1S2 := strings.Split(readLine(reader), " ")

	lTemp, err := strconv.ParseInt(lS1S2[0], 10, 64)
	checkError(err)
	l := int32(lTemp)

	s1Temp, err := strconv.ParseInt(lS1S2[1], 10, 64)
	checkError(err)
	s1 := int32(s1Temp)

	s2Temp, err := strconv.ParseInt(lS1S2[2], 10, 64)
	checkError(err)
	s2 := int32(s2Temp)

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int64

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := queriesItemTemp
		queries = append(queries, queriesItem)
	}

	result := movingTiles(l, s1, s2, queries)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%.20f", resultItem)

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
