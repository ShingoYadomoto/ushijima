package helper

import (
	"strconv"
)

func AtoiSlice(stringSlice []string) (intSlice []int, err error) {
	intSlice = make([]int, len(stringSlice))

	for i, v := range stringSlice {
		var value int
		value, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		intSlice[i] = value
	}
	return
}

func Atoi64(s string) (i int64, err error) {
	return strconv.ParseInt(s, 10, 64)
}
