package arr

// Remove Duplicates from Sorted Array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}
