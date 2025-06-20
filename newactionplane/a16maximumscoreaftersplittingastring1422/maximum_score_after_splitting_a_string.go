package a16maximumscoreaftersplittingastring1422

import (
	"strings"
)

func maxScore(s string) int {
	if len(s) == 1 {
		return 0
	}

	total1 := 0
	for _, v := range s {
		if v == '1' {
			total1++
		}
	}

	num0, num1, max := 0, 0, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			num1++
		}
		if s[i] == '0' {
			num0++
		}
		tmp := num0 + (total1 - num1)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// 问题：分割字符串，统计左边0和右边1和值的最大值
// 方法一:
// 设分割线位于s[i]和s[i+1]之间，分割线从-1开始移动
// 初始化：初始分割线位于-1，统计右边1的值
// 遍历移动：从0开始移动分割线，s[i]为0 ，左串+1，右串不变；s[i]为1，左串不变，右串-1，记录最大值，至n-2结束
func maxScore1(s string) int {
	right1 := strings.Count(s, "1") // 分割线位于-1，右串1的个数
	left0, ans := 0, 0
	for _, c := range s[:len(s)-1] {
		if c == '0' {
			left0++ // 分割线左边0的个数
		} else {
			right1-- // 分割线右边1的个数
		}
		ans = max(ans, left0+right1)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 方法二：将left0和right1合成一个变量score
// 遍历分割线，若遇到0，左串+1，右串不变，总量+1；若遇到1，左串不变，右串-1，总量-1
func maxScore2(s string) int {
	score := strings.Count(s, "1")
	ans := 0
	for _, v := range s[:len(s)-1] {
		if v == '0' {
			score++
		} else if v == '1' {
			score--
		}
		ans = max(ans, score)
	}
	return ans
}

// 方法3：TODO
// 设total0为串s总的0的个数，n为串s的总长度，分割线位于s[i]和s[i+1]之间，0=<i<=n-2
// 移动分割线，设left0为左串中0的个数，则right1=n-(i+1)-(total0-left0)
// score = left0 + n-(i+1)-(total0-left0) = (2 * left0 - i) + (n-1-total0)，
// 求score的最大值 => 求 2 * left0 - i 的最大值
func maxScore3(s string) int {
	left0, maxPart, n := 0, 0, len(s)
	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			left0++
		}
		tmp := 2*left0 - i
		if tmp > maxPart {
			maxPart = tmp
		}
	}
	if n >= 2 && s[n-1] == '0' {
		left0++
	}
	return maxPart + n - 1 - left0
}

// 方法4：maxScore = max(TotalLeftZeros + TotalRightOnes)
//                = max(TotalLeftZeros + TotalOnes - TotalLeftOnes)
//                = max(TotalLeftZeros - TotalLeftOnes) + TotalOnes
// 设分割线位于s[i]和s[i+1]之间，分割线从0开始移动，i位于[0, n-2]
// 最后判断s[n-1]是否为1，保证TotalOnes的正确性

func maxScore4(s string) int {
	left0, left1, maxPart, n := 0, 0, -1, len(s)
	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			left0++
		} else {
			left1++
		}
		maxPart = max(left0-left1, maxPart)
	}
	if s[n-1] == '1' {
		left1++
	}
	return maxPart + left1
}
