package arr

func searchInsert(nums []int, target int) int {
	for k, num := range nums {
		if target <= num {
			return k
		}
	}
	return len(nums)
}
