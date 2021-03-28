package arr

// 20200722-TX
// Input: [1,0,2,3,0,4,5,0]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

func duplicateZeros(arr []int) {
	if len(arr) <= 1 {
		return
	}
	//j:=len(arr)-1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j != i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
}

func duplicateZeros2(arr []int) {
	// 1. 计算需要复写0的数量
	pDups := 0
	length := len(arr) - 1
	for left := 0; left <= length-pDups; left++ {
		if arr[left] == 0 {
			// 2. 注意处理元素边界上0的情况
			if left == length-pDups {
				arr[length] = 0
				length--
				break
			}
			pDups++
		}
	}

	// 3. 从末尾迭代数组
	last := length - pDups
	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+pDups] = 0
			pDups--
			arr[i+pDups] = 0
		} else {
			arr[i+pDups] = arr[i]
		}
	}
}
