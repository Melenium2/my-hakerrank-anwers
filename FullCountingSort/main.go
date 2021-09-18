package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convInt(number string) int {
	temp, _ := strconv.ParseInt(number, 10, 32)
	return int(temp)
}

func countSort(arr [][]string) {

	//for i := 0; i < len(arr); i++ {
	//
	//}
	//
	//fmt.Fprint(os.Stdout, strings.Join(out, ""))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	out := make([]string, 100)
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

		var arrRow []string
		for _, arrRowItem := range arrRowTemp {
			arrRow = append(arrRow, arrRowItem)
		}

		if len(arrRow) != 2 {
			panic("Bad input")
		}

		k := convInt(arrRow[0])
		if i < int(n)/2 {
			out[k] += "-" + " "
		} else {
			out[k] += arrRow[1] + " "
		}
	}

	fmt.Fprint(os.Stdout, strings.Join(out, ""))
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
