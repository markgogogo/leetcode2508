package a10subtractproductandsum1281

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	if n == 0 {
		return 0
	}
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += digit
		product *= digit
	}
	return product - sum
}
