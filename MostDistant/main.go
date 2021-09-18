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

func dist(x, y float64) float64 {
	sum := x * x + y * y
	if sum > 0 {
		return math.Sqrt(sum)
	} else {
		return 0
	}
}

func max(s []int32) int32 {
	var max int32 = 0
	for _, v := range s {
		max = int32(math.Max(float64(max), float64(v)))
	}

	return max
}

func maxF(s []float64) float64 {
	var max float64 = 0
	for _, v := range s {
		max = math.Max(max, v)
	}

	return max
}

func min(s []int32) int32 {
	var min int32 = 0
	for _, v := range s {
		min = int32(math.Min(float64(min), float64(v)))
	}

	return min
}

func solve(c [][]int32) float64 {
	var xs []int32
	var ys []int32
	for _, v := range c {
		xs = append(xs, v[0])
		ys = append(ys, v[1])
	}

	xMax := max(xs)
	xMin := min(xs)
	yMax := max(ys)
	yMin := min(ys)
	x1 := math.Max(float64(xMax), math.Abs(float64(xMin)))
	x2 := math.Max(float64(yMax), math.Abs(float64(yMin)))
	r := math.Max(math.Max(float64(xMax - xMin), float64(yMax - yMin)), dist(x1, x2))

	return r
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var coordinates [][]int32
	for coordinatesRowItr := 0; coordinatesRowItr < int(n); coordinatesRowItr++ {
		coordinatesRowTemp := strings.Split(readLine(reader), " ")

		var coordinatesRow []int32
		for _, coordinatesRowItem := range coordinatesRowTemp {
			coordinatesItemTemp, err := strconv.ParseInt(coordinatesRowItem, 10, 64)
			checkError(err)
			coordinatesItem := int32(coordinatesItemTemp)
			coordinatesRow = append(coordinatesRow, coordinatesItem)
		}

		if len(coordinatesRow) != int(2) {
			panic("Bad input")
		}

		coordinates = append(coordinates, coordinatesRow)
	}

	result := solve(coordinates)

	fmt.Fprintf(writer, "%.10f\n", result)

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
