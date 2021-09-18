package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var modulus int64 = 1

type ccomplex struct {
	base      int64
	imaginary int64
}

func mul(first, second ccomplex) ccomplex {
	f := (first.base * second.base + modulus - first.imaginary * second.imaginary) % modulus
	for f < 0 {
		f += modulus
	}
	s := (first.base * second.imaginary + first.imaginary * second.base) % modulus

	return ccomplex{f, s}
}

func binpow(c ccomplex, exp int64) ccomplex {
	x := ccomplex{
		base: 1,
		imaginary: 0,
	}
	for exp > 0 {
		if exp&1 == 1 {
			x = mul(x, c)
		}
		c = mul(c, c)
		exp >>= 1
	}

	return x
}

func solve(a int64, b int64, k int64, m int64) []int64 {
	modulus = m
	c := binpow(ccomplex{
		base:      a,
		imaginary: b,
	}, k)

	return []int64{(c.base) % modulus , (c.imaginary) % modulus}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		abkm := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(abkm[0], 10, 64)
		checkError(err)
		a := aTemp

		bTemp, err := strconv.ParseInt(abkm[1], 10, 64)
		checkError(err)
		b := bTemp

		kTemp, err := strconv.ParseInt(abkm[2], 10, 64)
		checkError(err)
		k := kTemp

		mTemp, err := strconv.ParseInt(abkm[3], 10, 64)
		checkError(err)
		m := mTemp

		result := solve(a, b, k, m)

		for resultItr, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)
			if resultItr != len(result)-1 {
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
