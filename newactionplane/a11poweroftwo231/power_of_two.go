package a11poweroftwo231

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 0 {
		m := n % 2
		n = n / 2
		if n > 0 && m != 0 {
			return false
		}
	}
	return true
}

func isPowerOfTwo1(n int) bool {
	return true
}
