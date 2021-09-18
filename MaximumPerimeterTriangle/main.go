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

func isPossible(a, b, c int64) bool {
	if a + b > c {
		return true
	}

	return false
}

func maximumPerimeterTriangle(sticks []int32) []int32 {
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] < sticks[j]
	})

	var maxP int64 = 0
	var answer []int32
	for i := 0; i < len(sticks)-2; i++ {
		a, b, c := int64(sticks[i]), int64(sticks[i+1]), int64(sticks[i+2])
		sum := a + b + c
		if isPossible(a, b, c) && sum > maxP {
			maxP = sum
			answer = []int32{int32(a), int32(b), int32(c)}
		}
	}

	if len(answer) < 3 {
		return []int32{-1}
	} else {
		return answer
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sticksTemp := strings.Split(readLine(reader), " ")

	var sticks []int32

	for i := 0; i < int(n); i++ {
		sticksItemTemp, err := strconv.ParseInt(sticksTemp[i], 10, 64)
		checkError(err)
		sticksItem := int32(sticksItemTemp)
		sticks = append(sticks, sticksItem)
	}

	result := maximumPerimeterTriangle(sticks)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
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
