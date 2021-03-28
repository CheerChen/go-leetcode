package easy

func twoSum(nums []int, target int) []int {
	numdict := make(map[int]int)
	for i, num := range nums {
		try := target - num
		if j, ok := numdict[try]; ok {
			if j == i {
				continue
			}
			return []int{j, i}
		}
		numdict[num] = i
	}
	return []int{}
}
