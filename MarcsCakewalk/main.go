package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func marcsCakewalk(calorie []int32) int64 {
	sort.Slice(calorie, func(i, j int) bool {
		return calorie[i] > calorie[j]
	})
	sum := 0
	for i, k := range calorie {
		subsum := math.Pow(2, float64(i)) * float64(k)
		sum += int(subsum)
	}

	return int64(sum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	calorieTemp := strings.Split(readLine(reader), " ")

	var calorie []int32

	for i := 0; i < int(n); i++ {
		calorieItemTemp, err := strconv.ParseInt(calorieTemp[i], 10, 64)
		checkError(err)
		calorieItem := int32(calorieItemTemp)
		calorie = append(calorie, calorieItem)
	}

	result := marcsCakewalk(calorie)

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
