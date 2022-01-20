package easy

import "strconv"

func STI(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}
