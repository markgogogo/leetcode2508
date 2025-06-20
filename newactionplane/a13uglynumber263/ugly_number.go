package a13uglynumber263

// 方法1.
// 根据丑数定义，0 和 负整数 一定不是丑数
// 当 n > 0 时，若 n 是丑数，则 n 可以写成 n=2^a×3^b×5^c的形式，
// 其中 a,b,c 都是非负整数。特别地，当 a,b,c 都是 0 时，n=1
// 为判断 n 是否满足上述形式，可以对 n 反复除以 2,3,5，直到 n 不再包含质因数 2,3,5。若剩下的数等于 1，则说明 n 不包含其他质因数，是丑数；否则，说明 n 包含其他质因数，不是丑数。
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for {
		if n%2 == 0 {
			n /= 2
			continue
		}
		if n%3 == 0 {
			n /= 3
			continue
		}
		if n%5 == 0 {
			n /= 5
			continue
		}
		break
	}

	return n == 1
}

// 方法2.
// 根据丑数定义，当 n > 0 时，若 n 是丑数，则 n 可以写成 n=2^a×3^b×5^c的形式
// 可以先将n反复除以3和5，若 n 是丑数，则n可以写成n=2^a，再将n是否丑数转换成判断是否是2的幂
func isUgly2(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n&(n-1) == 0
}
