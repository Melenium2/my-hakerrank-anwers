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

func gridlandMetro(n int64, m int64, k int64, track [][]int64) int64 {
	sort.SliceStable(track, func(i, j int) bool {
		return track[i][0] < track[j][0]
	})

	var counter int64 = n * m
	var row, start, end int64 = 0, 0, 0
	for _, t := range track {
		r, s, e := t[0], t[1], t[2]
		if row == r && s - 1 < end {
			end = int64(math.Max(float64(end), float64(e)))
		} else {
			counter -= end - start
			row, start, end = r, s-1, e
		}
	}


	return counter - end + start
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nmk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nmk[0], 10, 64)
	checkError(err)
	n := int64(nTemp)

	mTemp, err := strconv.ParseInt(nmk[1], 10, 64)
	checkError(err)
	m := int64(mTemp)

	kTemp, err := strconv.ParseInt(nmk[2], 10, 64)
	checkError(err)
	k := int64(kTemp)

	var track [][]int64
	for i := 0; i < int(k); i++ {
		trackRowTemp := strings.Split(readLine(reader), " ")

		var trackRow []int64
		for _, trackRowItem := range trackRowTemp {
			trackItemTemp, err := strconv.ParseInt(trackRowItem, 10, 64)
			checkError(err)
			trackItem := int64(trackItemTemp)
			trackRow = append(trackRow, trackItem)
		}

		if len(trackRow) != int(3) {
			panic("Bad input")
		}

		track = append(track, trackRow)
	}

	result := gridlandMetro(n, m, k, track)

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
