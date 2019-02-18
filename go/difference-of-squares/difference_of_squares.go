package diffsquares

import (
	"math"
)

// SquareOfSum ...
func SquareOfSum(n int) int {

	result := 0
	for i := 0; i <= n; i++ {
		result = result + i
	}
	return int(math.Pow(float64(result), 2))
}

// SumOfSquares ...
func SumOfSquares(n int) int {

	result := 0
	for i := 0; i <= n; i++ {
		result = result + int(math.Pow(float64(i), 2))
	}
	return result
}

// Difference ...
func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)

}

/*
$ go test -v --bench . --benchmem
=== RUN   TestSquareOfSum
--- PASS: TestSquareOfSum (0.00s)
=== RUN   TestSumOfSquares
--- PASS: TestSumOfSquares (0.00s)
=== RUN   TestDifference
--- PASS: TestDifference (0.00s)
goos: linux
goarch: amd64
pkg: github.com/pbrao/exercism/go/difference-of-squares
BenchmarkSquareOfSum-4          20000000                77.4 ns/op             0 B/op             0 allocs/op
BenchmarkSumOfSquares-4           500000              3898 ns/op               0 B/op             0 allocs/op
BenchmarkDifference-4             300000              3544 ns/op               0 B/op             0 allocs/op
PASS
ok      github.com/pbrao/exercism/go/difference-of-squares      4.758s
*/
