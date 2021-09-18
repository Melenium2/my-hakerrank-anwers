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
func gcd(a, b int32) int32 {
	var i int32
	var r int32 = 1
	for i = int32(math.Max(float64(a), float64(b))); i > 1 ; i-- {
		if a%i == 0 && b%i == 0{
			r = i
			break
		}
	}

	return r
}

func restaurant(l int32, b int32) int32 {
	sq := gcd(l, b)
	return (l*b) / (sq*sq)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		lb := strings.Split(readLine(reader), " ")

		lTemp, err := strconv.ParseInt(lb[0], 10, 64)
		checkError(err)
		l := int32(lTemp)

		bTemp, err := strconv.ParseInt(lb[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := restaurant(l, b)

		fmt.Fprintf(writer, "%d\n", result)
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
