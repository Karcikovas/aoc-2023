package utils

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func Combinations(n, k int) int {
	if k > n {
		return 0
	}

	return factorial(n) / (factorial(k) * factorial(n-k))
}
