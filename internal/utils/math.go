package utils

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of a and b.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GCD returns the greatest common divisor of a and b using Euclid's algorithm.
func GCD(a, b int) int {
	// Use absolute values to handle negative numbers
	a, b = Abs(a), Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int {
	// Handle zero inputs
	if a == 0 || b == 0 {
		return 0
	}
	
	// Use absolute values for calculation
	absA, absB := Abs(a), Abs(b)
	gcd := GCD(absA, absB)
	
	// Divide first to reduce overflow risk
	return absA / gcd * absB
}

// Sign returns 1 for positive, -1 for negative, and 0 for zero.
func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// Sum returns the sum of all numbers in the slice.
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Product returns the product of all numbers in the slice.
// Returns 0 for empty slices.
func Product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}

// MinMax returns both the minimum and maximum values in the slice.
// Returns (0, 0) for empty slices.
func MinMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
