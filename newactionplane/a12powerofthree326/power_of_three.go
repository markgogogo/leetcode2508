package a12powerofthree326

func isPowerOfThree(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		m := n % 3
		if m != 0 {
			return false
		}
		n /= 3
	}
	return false
}

// 方法1. 试除法：n = 3^a => 将n反复整除3，直至不可整除，判断结果是否为1
func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

// 
func isPowerOfThree2(n int) bool {
	return true
}