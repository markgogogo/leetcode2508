package a01addtwointegers2235

func sum3(num1 int, num2 int) int { // 求“不进位和与进位和”之和，直至进位和清0 // TODO: 尾递归？
	if num2 == 0 {
		return num1
	}
	return sum3(num1^num2, (num1&num2)<<1)
}
