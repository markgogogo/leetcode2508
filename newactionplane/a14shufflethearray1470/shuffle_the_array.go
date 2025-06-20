package a14shufflethearray

// 方法1. 新数组（遍历原数组视角）：前n个元素(i从0~n-1)放到2*i位置；后n个元素(i从0~n-1, j从n到2n-1)放到2*i+1位置
func shuffle1(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
	}
	for i, j := 0, n; j < 2*n; i, j = i+1, j+1 {
		res[2*i+1] = nums[j]
	}
	return res
}

func shuffle1_1(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[n+i]
	}
	return res
}

// 方法2. 新数组（遍历原数组视角）：
// 最终数据为nums[0],nums[n],nums[1],nums[n+1]...nums[n-1],nums[2*n-1]
// 遍历前n个元素（i从0~n-1），新数组每次赋值两个元素nums[i]和nums[n+i]
func shuffle2(nums []int, n int) []int {
	res := make([]int, 2*n)
	index := 0
	for i := 0; i < n; i++ {
		res[index] = nums[i]
		index++
		res[index] = nums[n+i]
		index++
	}
	return res
}

// 方法3. 新数组（遍历新数组视角）：遍历新数组（j从0~2*n-1，i从0~n-1），j为偶数，从前n个数里取，j为奇数，从后n个数里取
func shuffle3(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i, j := 0, 0; j <= 2*n-1; i, j = i+1, j+1 {
		if j%2 == 0 {
			res[j] = nums[i]
		} else {
			res[j] = nums[n+i]
		}
	}
	return res
}

// 方法4. 新数组，遍历老数组，找到每个元素对应的新数组位置，赋值新数组
// i < n => 新位置：2 * i
// i >= n => 新位置：2 * j + 1 => n （j从0~n-1） => 2*(i-n) + 1
// 方法A1：新数组
func shuffle4(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		j := 0
		if i < n {
			j = 2 * i
		} else {
			j = 2*(i-n) + 1 // ? 如何有效推出来？
		}
		res[j] = nums[i]
	}
	return res
}

// 方法5. 原地更新，位运算，充分利用数值范围：1 <= nums[i] <= 10^3 => nums[i]最多占用后10位，可利用其他位数来暂存新赋值，最后再轮询用新赋值替换旧值
// 遍历数组，nums[i] = 新赋值 << 10 | nums[i]; 新赋值 = nums[j] && 1024 (新赋值可能已经做过位处理)
func shuffle5(nums []int, n int) []int {
	index, mask := 0, (1<<10)-1
	for i := 0; i < n; i++ {
		nums[index] = ((nums[i] & mask) << 10) | nums[index]
		index++
		nums[index] = ((nums[n+i] & mask) << 10) | nums[index]
		index++
	}
	for i := 0; i < 2*n; i++ {
		nums[i] = nums[i] >> 10
	}
	return nums
}

// 方法A2：原地更新，位运算
func shuffle5_1(nums []int, n int) []int {
	mask := (1 << 10) - 1
	for i, j := 0, 0; i < 2*n; i++ {
		if i < n {
			j = 2 * i
		} else {
			j = (i-n)*2 + 1
		}

		nums[j] = ((nums[i] & mask) << 10) | nums[j]
	}
	for i := 0; i < 2*n; i++ {
		nums[i] = nums[i] >> 10
	}
	return nums
}

// 位运算，集中在前n个元素
// 新序列：nums[0],nums[n],nums[1],nums[n+1],...,nums[n-1],nums[2*n-1]
// => 新序列是nums[i]和nums[n+i]构成一对，i从0到n-1
// => 利用nums元素<1000的条件，将nums[i]和nums[n+i]的值放到一起，nums[i]占用低10位，nums[n+i]占用中间的10位
// 再将合并好的前n个元素，分别放出来，i对应到2*i+1和2*i位置
func shuffle5_2(nums []int, n int) []int {
	for i := 0; i <= n-1; i++ {
		nums[i] |= (nums[n+i] << 10)
	}
	for i := n - 1; i >= 0; i-- { // 从后往前追，两两结对赋值时，先赋后值。否则，nums[0]有两两合并结果，先赋0，会导致nums[0]值变更，导致合并结果丢失，nums[1]从nums[0]拿不到正确的数
		nums[2*i+1] = nums[i] >> 10
		nums[2*i] = nums[i] & 1023
	}
	return nums
}

// 方法6. 单变量 todo todo2 todo3
// 通过in-place swap的方法做到O(1)空间O(n)时间。
// 每个"nums[i]"都有一个“目标”index。例如对于8个数的nums, "nums[0]"想去"0", "nums[4]"想去"1", "nums[1]"想去"2", "nums[5]"想去"3", "nums[2]"想去"4"...
// in-place把nums[i] swap到它想去的index，把swap走的数标记为负数，并把swap回来的数继续"原地"swap出去，直到swap回来的数的目标index就是“i”自己，然后才增加"i"并继续过下一个“i”。
// 遇到nums[i]是负数就说明nums[i]已经在之前的swap中到达了目前位置，因此跳过。
// 所有的i都过好后nums就是正确的顺序，别忘了最后再过一遍把所有的负数变回正数。
// 由于每个nums[i]只会被标记1次负数，因此时间复杂度是O(n)
func shuffle6(nums []int, n int) []int {
	for i := 0; i < 2*n; i++ {
		if nums[i] < 0 {
			continue
		}
		// i: 当前元素nums[i]的索引，保持不变，标识是否终止
		// j: 每轮交换后的索引，用于计算下一轮的索引
		// k: 下一轮的索引
		j, k := i, 0
		for {
			if j < n {
				k = 2 * j
			} else {
				k = (j-n)*2 + 1
			}
			if i == k {
				nums[i] = -1 * nums[i]
				// fmt.Printf("i:%d, j:%d, k:%d, nums[i]:%d, nums[j]: %d, nums[k]:%d\n", i, j, k, nums[i], nums[j], nums[k])
				break
			} else {
				nums[i], nums[k] = nums[k], nums[i]*(-1)
				// fmt.Printf("i:%d, j:%d, k:%d, nums[i]:%d, nums[j]: %d, nums[k]:%d\n", i, j, k, nums[i], nums[j], nums[k])
				j = k
			}
		}
	}
	for i := 0; i < 2*n; i++ {
		nums[i] = -1 * nums[i]
	}
	return nums
}

// 方法A3：原地数组。利用每个元素>0的特点，用负数标识元素已在正确的位置
// 每个元素均有其置换完后的正确位置。从任一元素出发，可以构成链式置换
// 如果元素为正数，则启动链式置换，将当前元素*负1和正确位置元素互换，用新下标继续计算下一位置下标，重复，直至当前元素为负数。
// 继续下一轮，至均置为负数
func shuffle6_1(nums []int, n int) []int {
	for i := 0; i < 2*n; i++ {
		if nums[i] > 0 {
			j := i // j描述当前的nums[i]对应的索引
			for nums[i] > 0 {
				// 计算j索引的元素
				if j < n {
					j = 2 * j
				} else {
					j = (j-n)*2 + 1
				}
				// 将 nums[i] 放置到j的位置，同时标记上负数，表示放到了正确的位置
				// 同时，将nums[j]的元素放到i的位置，继续下一轮循环，直至i的位置放置了正确元素
				nums[i], nums[j] = nums[j], nums[i]*(-1)
			}
		}
	}

	for i := 0; i < 2*n; i++ {
		nums[i] = -1 * nums[i]
	}
	return nums
}
