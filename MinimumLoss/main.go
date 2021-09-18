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

func partition(arr []int64, l, r int) int {
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

func quickSort(arr []int64, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
}

// Complete the minimumLoss function below.
func minimumLoss(price []int64) int64 {
	initial := make(map[int64]int)
	for i, n := range price {
		initial[n] = i
	}
	quickSort(price, 0, len(price) - 1)

	diff := int64(math.MaxInt64)

	for i := 0; i < len(price)-1; i++ {
		d := price[i+1] - price[i]
		if d < diff && initial[price[i+1]] < initial[price[i]] {
			diff = d
		}
	}

	return diff
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024 * 16)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024 * 16)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := nTemp

	priceTemp := strings.Split(readLine(reader), " ")

	var price []int64

	for i := 0; i < int(n); i++ {
		priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
		checkError(err)
		price = append(price, priceItem)
	}

	result := minimumLoss(price)

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
