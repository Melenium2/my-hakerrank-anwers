package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int32, l, r int) int {
	pivot := arr[r]
	i := l

	for j := l; j <= r; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[r] = arr[r], arr[i]

	return i
}

func quickSort(arr []int32, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
}

func chunks(x []int32, k int) [][]int32 {
	n := k*2+1
	var chunks [][]int32
	for i := 0; i < len(x);  {
		ch := i+int(n)
		if ch > len(x) {
			ch = len(x)
		}
		chunks = append(chunks, x[i:ch])
		i += ch
	}

	return chunks
}

func hackerlandRadioTransmitters(x []int32, k int32) int {
	quickSort(x, 0, len(x) - 1)

	counter := 0
	var cover, min int32 = 0, 0
	for i := 0; i < len(x); i++ {
		if cover < x[i] {
			min = x[i]
			counter++
			cover = min+k
		} else if x[i] - min <= k {
			cover = x[i] + k
		}
	}

	return counter
}


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	xTemp := strings.Split(readLine(reader), " ")

	var x []int32

	for i := 0; i < int(n); i++ {
		xItemTemp, err := strconv.ParseInt(xTemp[i], 10, 64)
		checkError(err)
		xItem := int32(xItemTemp)
		x = append(x, xItem)
	}

	result := hackerlandRadioTransmitters(x, k)

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
