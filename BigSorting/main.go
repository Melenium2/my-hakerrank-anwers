package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func compare(f, l string) bool {
	lenF := len(f)
	lenL := len(l)
	if lenF != lenL { return lenF < lenL }

	for i := 0; i < lenF; i++ {
		if f[i] != l[i] {
			return f[i] < l[i]
		}
	}

	return false
}

func mergeSort(sl []string) []string {
	s := len(sl)
	if s == 1 {
		return sl
	}

	middle := s / 2
	var (
		left = make([]string, middle)
		right = make([]string, s - middle)
	)
	for i := 0; i < s; i++ {
		if i < middle {
			left[i] = sl[i]
		} else {
			right[i-middle] = sl[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(l, r []string) []string {
	result := make([]string, len(l) + len(r))

	i := 0
	for len(l) > 0 && len(r) > 0 {
		if compare(l[0], r[0]) {
			result[i] = l[0]
			l = l[1:]
		} else {
			result[i] = r[0]
			r = r[1:]
		}
		i++
	}

	for j := 0; j < len(l); j++ {
		result[i] = l[j]
		i++
	}
	for j := 0; j < len(r); j++ {
		result[i] = r[j]
		i++
	}

	return result
}

// Complete the bigSorting function below.
func bigSorting(unsorted []string) []string {
	return mergeSort(unsorted)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var unsorted []string

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		unsorted = append(unsorted, unsortedItem)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result) - 1 {
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
