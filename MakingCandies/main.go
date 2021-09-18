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

func minimumPasses(m int64, w int64, p int64, n int64) int64 {
	if m*w <= 0 {
		return 1
	}

	var days, canides, answer int64 = 0, 0, int64(math.Ceil(float64(n) / float64(m*w)))

	for days < answer {
		if p > canides {
			d := int64(math.Ceil((float64(p - canides)) / (float64(m * w))))
			canides += d * m * w
			days += d
		}

		diff := int64(math.Abs(float64(m - w)))
		var available int64 = 0
		available = canides / p
		buy := int64(math.Min(float64(diff), float64(available)))

		if m < w {
			m += buy
		} else {
			w += buy
		}

		rest := available - buy
		m += rest/2
		w += rest - rest/2
		canides -= available * p


		canides += m * w
		days += 1

		if canides > n {
			days++
			return int64(math.Abs(float64(days-answer)))
		}

		remain := int64(math.Max(float64(n - canides), 0))

		answer = int64(math.Min(float64(answer), float64(days) + math.Ceil(float64(remain) / float64(m * w))))
	}

	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	mwpn := strings.Split(readLine(reader), " ")

	m, err := strconv.ParseInt(mwpn[0], 10, 64)
	checkError(err)

	w, err := strconv.ParseInt(mwpn[1], 10, 64)
	checkError(err)

	p, err := strconv.ParseInt(mwpn[2], 10, 64)
	checkError(err)

	n, err := strconv.ParseInt(mwpn[3], 10, 64)
	checkError(err)

	result := minimumPasses(m, w, p, n)

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
