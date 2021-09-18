package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bestDivisor(n int) int {
	var i int
	max, answer := 0, 0
	for i = 1; i <= n; i++ {
		if n % i == 0 {
			dig := i
			sum := 0
			for dig != 0 {
				sum += dig%10
				dig = dig/10
			}
			if max < sum {
				max = sum
				answer = i
			}
		}
	}
	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	res := bestDivisor(n)

	fmt.Fprintf(os.Stdout, "%d\n", res)
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
