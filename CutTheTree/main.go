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

func dfs(root, parent, totalSum int32, edges [][]int32, subtree []int32, res *int32) {
	sum := subtree[root]

	for i := 0; i < len(edges[root]); i++ {
		v := edges[root][i]
		if v != parent {
			dfs(v, root, totalSum, edges, subtree, res)
			sum += subtree[v]
		}
	}

	subtree[root] = sum

	fmt.Fprintf(os.Stdout, "Root: %d, parent: %d, sum: %d\n", root, parent, sum)

	subsum := int32(math.Abs(float64(totalSum - 2*sum)))
	if root != 0 && subsum < *res {
		*res = subsum
	}
}

func sumDifference(vertex []int32, edges [][]int32, N int) int32 {
	var totalSum int32 = 0
	subtree := make([]int32, N)

	for i := 0; i < N; i++ {
		subtree[i] = vertex[i]
		totalSum += vertex[i]
	}

	edge := make([][]int32, N)
	for i := 0; i < N-1; i++ {
		edge[edges[i][0]-1] = append(edge[edges[i][0]-1], edges[i][1]-1)
		edge[edges[i][1]-1] = append(edge[edges[i][1]-1], edges[i][0]-1)
	}

	var res = int32(math.MaxInt32)

	dfs(0, -1, totalSum, edge, subtree, &res)
	return res
}

func cutTheTree(vertex []int32, edges [][]int32) int32 {
	return sumDifference(vertex, edges, len(vertex))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	dataTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var data []int32

	for i := 0; i < int(n); i++ {
		dataItemTemp, err := strconv.ParseInt(dataTemp[i], 10, 64)
		checkError(err)
		dataItem := int32(dataItemTemp)
		data = append(data, dataItem)
	}

	var edges [][]int32
	for i := 0; i < int(n) - 1; i++ {
		edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

		var edgesRow []int32
		for _, edgesRowItem := range edgesRowTemp {
			edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
			checkError(err)
			edgesItem := int32(edgesItemTemp)
			edgesRow = append(edgesRow, edgesItem)
		}

		if len(edgesRow) != 2 {
			panic("Bad input")
		}

		edges = append(edges, edgesRow)
	}

	result := cutTheTree(data, edges)

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
