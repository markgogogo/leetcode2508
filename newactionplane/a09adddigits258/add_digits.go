package a09adddigits258

func addDigits(num int) int {
	res := addRound(num)
	for res >= 10 {
		res = addRound(res)
	}
	return res
}

func addRound(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num = num / 10
	}
	return res
}
