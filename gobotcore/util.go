package gobotcore

import "strconv"

func AtoiEZPZ(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return -1
}
