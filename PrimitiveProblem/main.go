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

func isPrime(n int32) bool {
	if n <= 1 { return false }
	if n <= 3 { return true }

	if n % 2 == 0 || n % 3 == 0 { return false }

	var i int32 = 0
	for i = 5; i*i <= n; i = i + 6 {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
	}

	return true
}

func modpow(base, exp, mod int64) int64 {
	base = base % mod
	var res int64 = 1
	for exp > 0 {
		if exp & 1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}

	return res
}

func primeFactors(n int32) map[int32]int32 {
	primes := make(map[int32]int32)
	for n % 2 == 0 {
		primes[2] = 0
		n /= 2
	}

	var i int32 = 0
	for i = 3; i <= int32(math.Sqrt(float64(n))); i = i + 2 {
		for n % i == 0 {
			primes[i] = 0
			n /= i
		}
	}

	if n > 2 {
		primes[n] = 0
	}

	return primes
}

func rootPrimitive(n int32) (int32, int) {
	if !isPrime(n) {
		return -1, 0
	}
	phi := n - 1
	primes := primeFactors(phi)

	var i int32 = 0
	var minPirm int32 = 0
	for i = 2; i <= phi; i++ {
		flag := false
		for k, _ := range primes {
			pow := modpow(int64(i), int64(phi / k), int64(n))
			if pow == 1 {
				flag = true
				break
			}
		}

		if !flag {
			minPirm = i
			break
		}
	}

	for k, _ := range primes {
		phi = phi/k * (k-1)
	}

	if minPirm == 0 {
		return -1, 0
	} else {
		return minPirm, int(phi)
	}
}

func solve(n int32) (int32, int) {
	return rootPrimitive(n)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	n1, n2 := solve(p)
	fmt.Fprintf(os.Stdout, "%d %d\n", n1, n2)
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
