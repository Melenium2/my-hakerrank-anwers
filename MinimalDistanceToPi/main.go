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

var pi = math.Pi - 3

func solve(a, b int64) string {
	var i int64 = 0
	var min float64 = 1
	var de, nu int64 = 0, 0
	for i = a; i <= b; i++ {
		fi := float64(i)
		p := math.Floor(fi * pi)
		dif1 := pi - p/fi
		dif2 := -pi + (p+1)/fi
		if dif1 < min {
			min = dif1
			de = i
			nu = int64(p)
		} else if dif2 < min {
			min = dif2
			de = i
			nu = int64(p+1)
		}
	}
	return fmt.Sprintf("%d/%d", nu + 3 * de, de)
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1024*1024)
	numbs := strings.Split(readLine(in), " ")
	a, _ := strconv.ParseInt(numbs[0], 10, 64)
	b, _ := strconv.ParseInt(numbs[1], 10, 64)

	result := solve(a, b)
	fmt.Fprintf(os.Stdout, "%s\n", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}