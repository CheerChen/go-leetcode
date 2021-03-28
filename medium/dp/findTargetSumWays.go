package dp

// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数

import "strconv"

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make(map[string]int)
	return dpFindTargetSumWays(nums, 0, S, memo)
}

func dpFindTargetSumWays(nums []int, l, S int, memo map[string]int) int {
	if l == len(nums) {
		if S == 0 {
			return 1
		}
		return 0
	}
	key := strconv.Itoa(l) + "," + strconv.Itoa(S)
	if res, ok := memo[key]; ok {
		return res
	}
	result := dpFindTargetSumWays(nums, l+1, S-nums[l], memo) + dpFindTargetSumWays(nums, l+1, S+nums[l], memo)
	memo[key] = result
	return result
}
