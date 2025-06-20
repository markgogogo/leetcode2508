package a01addtwointegers2235

func sum2(num1 int, num2 int) int { // 求“不进位和与进位和”之和，进位和迭代左移直至清0 // TODO: 位计算，加减乘除的模拟
	for num2 != 0 {
		carray := (num1 & num2) << 1
		num1 ^= num2
		num2 = carray
	}

	return num1
}
