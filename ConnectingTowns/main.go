package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func connectingTowns(n int32, routes []int32) int32 {
	var multi int32 = 1
	var i int32 = 0
	for i < n-1 {
		multi = multi*routes[i]%1234567
		i++
	}

	return multi
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		routesTemp := strings.Split(readLine(reader), " ")

		var routes []int32

		for routesItr := 0; routesItr < int(n-1); routesItr++ {
			routesItemTemp, err := strconv.ParseInt(routesTemp[routesItr], 10, 64)
			checkError(err)
			routesItem := int32(routesItemTemp)
			routes = append(routes, routesItem)
		}

		result := connectingTowns(n, routes)

		fmt.Fprintf(os.Stdout, "%d\n", result)
	}
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
