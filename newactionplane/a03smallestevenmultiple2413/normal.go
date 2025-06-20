package a03smallestevenmultiple2413

func smallestEvenMultiple(n int) int {
	return (n%2 + 1) * n
}

func smallestEvenMultiple1(n int) int {
	if n%2 == 0 {
		return n
	}

	return 2 * n
}
