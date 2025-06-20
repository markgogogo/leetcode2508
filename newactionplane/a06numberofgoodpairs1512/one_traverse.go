package a06numberofgoodpairs1512

func numIdenticalPairs3(nums []int) int {
	res := 0
	numCnt := map[int]int{}
	for i := 0; i < len(nums); i++ {
		res += numCnt[nums[i]]
		numCnt[nums[i]]++
	}
	return res
}
