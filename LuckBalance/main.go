package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumAr(arr []int32) int32 {
	var sum int32 = 0
	for _, n := range arr {
		sum += n
	}

	return sum
}

// Complete the luckBalance function below.
func luckBalance(k int32, contests [][]int32) int32 {
	var sum int32 = 0
	important := make([]int32, 0)

	for i := 0; i < len(contests); i++ {
		score, isImprotant := contests[i][0], contests[i][1]
		sum += score
		if isImprotant == 1 {
			important = append(important, score)
		}
	}

	if k == int32(len(important)) || len(important) == 0 || int32(len(important)) < k {
		return sum
	} else if k == 0 {
		return sum - sumAr(important)*2
	} else {
		sort.Slice(important, func(i, j int) bool {
			return important[i] < important[j]
		})
		return sum - sumAr(important[0:int32(len(important))-k])*2
	}
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

	var contests [][]int32
	for i := 0; i < int(n); i++ {
		contestsRowTemp := strings.Split(readLine(reader), " ")

		var contestsRow []int32
		for _, contestsRowItem := range contestsRowTemp {
			contestsItemTemp, err := strconv.ParseInt(contestsRowItem, 10, 64)
			checkError(err)
			contestsItem := int32(contestsItemTemp)
			contestsRow = append(contestsRow, contestsItem)
		}

		if len(contestsRow) != 2 {
			panic("Bad input")
		}

		contests = append(contests, contestsRow)
	}

	result := luckBalance(k, contests)

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
