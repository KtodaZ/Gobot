package gobotcore

import "strconv"

func AtoiEZPZ(str string) int8 {
	if i, err := strconv.Atoi(str); err == nil {
		return int8(i)
	}
	return -1
}
