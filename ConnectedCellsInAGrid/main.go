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

const k int = 8
var rowsCube = [k]int{-1, -1, -1, 0, 0, 1, 1, 1 }
var colsCube = [k]int{-1, 0, 1, -1, 1, -1, 0, 1 }

func isConnected(i, j int, matrix [][]int32, visited [][]bool) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])  && !visited[i][j] && matrix[i][j] == 1
}

func dfs(i, j int, counter *int, matrix [][]int32, visited [][]bool) {
	visited[i][j] = true
	for z := 0; z < k; z++ {
		if isConnected(i + rowsCube[z], j + colsCube[z], matrix, visited) {
			*counter++
			dfs(i + rowsCube[z], j + colsCube[z], counter, matrix, visited)
		}
	}
}

// Complete the connectedCell function below.
func connectedCell(matrix [][]int32) int {
	sum := math.MinInt32
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			visited[i] = append(visited[i], false)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				c := 1
				dfs(i, j, &c, matrix, visited)
				sum = int(math.Max(float64(c), float64(sum)))
			}
		}
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var matrix [][]int32
	for i := 0; i < int(n); i++ {
		matrixRowTemp := strings.Split(readLine(reader), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(int(m)) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	result := connectedCell(matrix)

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
