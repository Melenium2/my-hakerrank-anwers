package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 3 5 5 7 8 11
// 5 5 7 8 10 11

func beautifulPairs(A []int32, B []int32) int32 {
	set := make(map[int32]int32)
	for _, n := range A {
		set[n]++
	}

	answer := 0
	for _, n := range B {
		v, ok := set[n]
		if ok && v > 0 {
			set[n]--
			answer++
		}
	}

	if answer == len(A) {
		answer--
	} else {
		answer++
	}

	return int32(answer)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	ATemp := strings.Split(readLine(reader), " ")

	var A []int32

	for i := 0; i < int(n); i++ {
		AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	BTemp := strings.Split(readLine(reader), " ")

	var B []int32

	for i := 0; i < int(n); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := beautifulPairs(A, B)

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

