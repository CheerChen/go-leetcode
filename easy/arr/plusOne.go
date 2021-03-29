package arr

func plusOne(digits []int) []int {
	res := []int{}
	i, carry := len(digits)-1, 0

	for i >= 0 {
		if i == len(digits)-1 {
			digits[i]++
		}
		if carry != 0 {
			digits[i]++
		}
		carry = digits[i] / 10
		res = append([]int{digits[i] % 10}, res...)
		i--
	}

	if carry != 0 {
		res = append([]int{1}, res...)
	}
	return res
}
