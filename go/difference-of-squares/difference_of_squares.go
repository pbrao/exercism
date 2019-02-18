package diffsquares

// SquareOfSum calculates the square of a sum of numbers
func SquareOfSum(n int) int {

	sum := n * (n + 1) / 2
	//return int(math.Pow(float64(sum), 2))
	return sum * sum
}

// SumOfSquares calculates the sum of a square of numbers
func SumOfSquares(n int) int {

	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between the square of sum and sum of squares
func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)

}
