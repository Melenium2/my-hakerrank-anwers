package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 2
// 1 2 3 4 5
//func partition(arr []int32, l, r int) int {
//	pivot := arr[r]
//	i := l
//
//	for j := l; j <= r; j++ {
//		if arr[j] < pivot {
//			arr[i], arr[j] = arr[j], arr[i]
//			i++
//		}
//	}
//
//	arr[i], arr[r] = arr[r], arr[i]
//
//	return i
//}
//
//func quicksort(arr []int32, l, r int) {
//	if l < r {
//		p := partition(arr, l, r)
//		quicksort(arr,l, p-1)
//		quicksort(arr, p+1, r)
//	}
//}

func pairs(k int32, arr []int32) int32 {
	m := make(map[int32]int)

	for _, n := range arr {
		m[n]++
	}

	countner := 0

	for i := 0; i < len(arr); i++ {
		countner += m[arr[i] - k]
	}

	return int32(countner)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024 * 15)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

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
