package arr

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			// 只是拷贝到前面到位置
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
