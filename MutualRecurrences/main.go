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

/**
 ХЗ как это рашать вообще
По сути тут необходимо сделать трансформирующию матрицу
Возвести матрицу в степень и домножить на вектор
 */

var (
	mod = int64(math.Pow(10, 9))
	N   = 30
)

type Matrix struct {
	size int
	m    [][]int64
}

func NewMatrix(size int) *Matrix {
	matr := make([][]int64, size)
	for i, _ := range matr {
		matr[i] = make([]int64, size)
	}

	return &Matrix{
		size: size,
		m: matr,
	}
}

func NewMatrixInitDi(size int, diagonal int64) *Matrix {
	matr := make([][]int64, size)
	for i, _ := range matr {
		matr[i] = make([]int64, size)
	}

	for i := 0; i < size; i++ {
		matr[i][i] = diagonal
	}

	return &Matrix{
		size: size,
		m: matr,
	}
}

func (mat Matrix) mul(matrix *Matrix) *Matrix {
	ans := NewMatrix(mat.size)
	for i := 0; i < mat.size-1; {
		i+=1
		for j := 0; j < mat.size-1; {
			j+=1
			for k := 0; k < mat.size-1; {
				k+=1
				ans.m[i][j] = add(ans.m[i][j], mul(mat.m[i][k], matrix.m[k][j]))
			}
		}
	}

	return ans
}

func (mat Matrix) mulByV(vector Vector) Vector {
	v := NewVector(24)
	for i := 0; i < 23; {
		i+=1
		for j := 0; j < 23; {
			j+=1
			v[i] = add(v[i], mul(mat.m[i][j], vector[j]))
		}
	}

	return v
}

type Vector []int64

func NewVector(size int) Vector {
	v := make([]int64, size)
	return v
}

func binpow(base *Matrix, exp int64) *Matrix {
	res := NewMatrixInitDi(N, 1)
	for exp > 0 {
		if exp&1 == 1 {
			res = res.mul(base)
		}
		base = base.mul(base)
		exp >>= 1
	}

	return res
}

func add(x, y int64) int64 {
	x += y
	if x >= mod {
		x -= mod
	}

	return x
}

func mul(x, y int64) int64 {
	return (1 * x * y)%mod
}

func solve(a int32, b int32, c int32, d int32, e int32, f int32, g int32, h int32, n int64) []int64 {
	m := NewMatrix(N)
	m.m[0][a - 1]++
	m.m[0][10 + b - 1]++
	m.m[0][10 + c - 1]++
	m.m[0][22]++

	m.m[10][10 + e - 1]++
	m.m[10][f - 1]++
	m.m[10][g - 1]++
	m.m[10][20]++

	for i := 1; i <= 9; {
		i+=1
		m.m[i][i-1]++
		m.m[i + 10][10 + i - 1]++
	}

	m.m[20][20] = int64(h)
	m.m[20][21] = int64(h)
	m.m[21][21] = int64(h)

	m.m[22][22] = int64(d)
	m.m[22][23] = int64(d)
	m.m[23][23] = int64(d)

	v := NewVector(24)
	v[20] = 0
	v[21] = 1
	v[22] = 0
	v[23] = 1

	m = binpow(m, n+1)

	v =	m.mulByV(v)

	return []int64{v[0], v[10]}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		abcdefghn := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(abcdefghn[0], 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(abcdefghn[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		cTemp, err := strconv.ParseInt(abcdefghn[2], 10, 64)
		checkError(err)
		c := int32(cTemp)

		dTemp, err := strconv.ParseInt(abcdefghn[3], 10, 64)
		checkError(err)
		d := int32(dTemp)

		eTemp, err := strconv.ParseInt(abcdefghn[4], 10, 64)
		checkError(err)
		e := int32(eTemp)

		fTemp, err := strconv.ParseInt(abcdefghn[5], 10, 64)
		checkError(err)
		f := int32(fTemp)

		gTemp, err := strconv.ParseInt(abcdefghn[6], 10, 64)
		checkError(err)
		g := int32(gTemp)

		hTemp, err := strconv.ParseInt(abcdefghn[7], 10, 64)
		checkError(err)
		h := int32(hTemp)

		n, err := strconv.ParseInt(abcdefghn[8], 10, 64)
		checkError(err)

		result := solve(a, b, c, d, e, f, g, h, n)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
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
