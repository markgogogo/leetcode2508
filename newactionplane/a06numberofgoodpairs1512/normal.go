// https://leetcode.cn/problems/number-of-good-pairs/?envType=study-plan-v2&envId=primers-list
package a06numberofgoodpairs1512

func numIdenticalPairs(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				cnt++
			}
		}
	}
	return cnt
}
