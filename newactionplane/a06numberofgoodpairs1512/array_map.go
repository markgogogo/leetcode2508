package a06numberofgoodpairs1512

func numIdenticalPairs2(nums []int) int {
	pairsCandidate := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pairsCandidate[nums[i]]++
	}
	res := 0
	for _, v := range pairsCandidate {
		res += factorial(v)
	}
	return res
}

func factorial(n int) int {
	return n * (n - 1) / 2
}
