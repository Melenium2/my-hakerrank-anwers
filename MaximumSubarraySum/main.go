package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var Mod int64 = 0
//
func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
//
//func maxOfThree(a, b, c int64) int64 {
//	return max(max(a, b), c)
//}
//
//func maxCrossesSum(arr []int64, l, r, m int) int64 {
//	var sum, leftSum, rightSum int64 = 0, 0, 0
//
//	for i := m; i >= l; i-- {
//		sum = (sum + arr[i]) % Mod
//		if sum > leftSum { leftSum = sum }
//	}
//
//	sum = 0
//	for i := m+1; i <= r; i++ {
//		sum = (sum + arr[i]) % Mod
//		if sum > rightSum { rightSum = sum }
//	}
//
//	return max(leftSum, rightSum) % Mod
//}
//
//func maxSubArray(arr []int64, l, r int) int64 {
//	if l == r {
//		return arr[l]
//	}
//
//	mid := (l + r) / 2
//
//	one := maxSubArray(arr, l, mid) % Mod
//	two := maxSubArray(arr, mid+1, r) % Mod
//	three := maxCrossesSum(arr, l, r, mid)
//
//	return maxOfThree(
//		one, two, three,
//	)
//}
//
//// Complete the maximumSum function below.
//func maximumSum(arr []int64, mod int64) int64 {
//	Mod = mod
//	return maxSubArray(arr, 0, len(arr) - 1)
//}

//func kadanesAlgo(arr []int64) int64 {
//	var local, global int64 = 0, 0
//
//	for i := 0; i < len(arr); i++ {
//		local = max(arr[i], arr[i] + local)
//		if local > global {
//			global = local
//		}
//	}
//
//	return global
//}

//func kadanesAlgo(arr []int64) int64 {
//	var local, global int64 = 0, 0
//
//	for i := 0; i < len(arr); i++ {
//		local += arr[i]
//		local %= Mod
//		if local > 0 {
//			global = max(global, local)
//		} else {
//			local = 0
//		}
//	}
//
//	return global
//}

func find(arr []int64) int64 {

	var pref, maxi int64 = 0, 0
	set := make(map[int64]int)

	for _, n := range arr {
		pref = (pref + n) % Mod

		maxi = max(maxi, pref)
		var v int64 = 0
		for k, _ := range set {
			if k >= pref {
				v = k
				break
			}
		}

		if v != 0 {
			maxi = max(maxi, pref - v + Mod)
		}
		set[pref]++
	}

	return maxi
}

// Complete the maximumSum function below.
func maximumSum(arr []int64, mod int64) int64 {
	Mod = mod
	return find(arr)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*15)

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		m, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)

		aTemp := strings.Split(readLine(reader), " ")

		var a []int64

		for i := 0; i < int(n); i++ {
			aItem, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			a = append(a, aItem)
		}

		result := maximumSum(a, m)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
