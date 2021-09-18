package main
//
//import (
//	"bufio"
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"strconv"
//	"strings"
//)
//
//
//
///*
// * Complete the 'getNumDraws' function below.
// *
// * The function is expected to return an INTEGER.
// * The function accepts INTEGER year as parameter.
// */
//
//const url = "https://jsonmock.hackerrank.com/api/football_matches?year=%d&page=%d&team1goals=%d&team2goals=%d"
//
//func getNumDraws(year int32) int32 {
//	page, count, goals := 1, 0, 0
//	for {
//		r, err := http.Get(fmt.Sprintf(url, year, page, goals, goals))
//		if err != nil {
//			return 0
//		}
//		obj := make(map[string]interface{})
//		if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
//			return 0
//		}
//		num := obj["total"].(float64)
//		if num == 0 {
//			break
//		}
//		count += int(num)
//
//		goals++
//	}
//
//	return int32(count)
//}
//
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
//
//	writer := bufio.NewWriterSize(os.Stdout, 16 * 1024 * 1024)
//
//	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	year := int32(yearTemp)
//
//	result := getNumDraws(year)
//
//	fmt.Fprintf(writer, "%d\n", result)
//
//	writer.Flush()
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
