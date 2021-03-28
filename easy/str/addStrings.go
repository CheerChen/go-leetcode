package str

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	res := ""
	i, j, carry := len(num1)-1, len(num2)-1, 0

	for i >= 0 || j >= 0 || carry > 0 {
		var s1 int
		if i >= 0 {
			s1, _ = strconv.Atoi(string(num1[i]))
		}

		var s2 int
		if j >= 0 {
			s2, _ = strconv.Atoi(string(num2[j]))
		}
		// fmt.Println(s1, s2)
		sum := s1 + s2
		if carry != 0 {
			sum++
		}
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
		i--
		j--
	}
	return res
}
