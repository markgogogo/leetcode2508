package a15transposematrix

func transpose(matrix [][]int) [][]int {
	// 1. 获取二维数组的行和列
	rows := len(matrix)
	if rows == 0 {
		return [][]int{}
	}
	cols := len(matrix[0])

	// 2. 新建二维数组
	ans := make([][]int, cols)
	for i := range ans {
		ans[i] = make([]int, rows)
	}

	// 3. 遍历二维数组
	for i, row := range matrix {
		for j, val := range row {
			ans[j][i] = val
		}
	}

	return ans
}
