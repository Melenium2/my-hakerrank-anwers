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

func primeFactors(n int32) []int32 {
	primes := make([]int32, 0)
	for n % 2 == 0 {
		primes = append(primes, 2)
		n /= 2
	}

	var i int32 = 0
	for i = 3; i < int32(math.Sqrt(float64(n))); i+=2 {
		for n % i == 0 {
			primes = append(primes, i)
			n /= i
		}
	}

	if n > 2 {
		primes = append(primes, n)
	}

	return primes
}

func digits(n int32) int32 {
	var digit []int32

	for n > 0 {
		v := n % 10
		digit = append(digit, v)
		n /= 10
	}

	var sum int32 = 0
	for _, d := range digit {
		sum += d
	}

	return sum
}

func solve(n int32) int32 {
	primes := primeFactors(n)
	var psum int32 = 0
	for _, v := range primes {
		psum += digits(v)
	}
	sum := digits(n)
	if sum == psum {
		return 1
	}
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := solve(n)

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
