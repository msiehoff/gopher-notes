package performance

import (
	"fmt"
	"strconv"
)

// MyFunc - func to benchmark
func MyFunc(ints []int) []string {
	return faster(ints)
	//return slower(ints)
}

func slower(ints []int) []string {
	s := []string{}
	for _, i := range ints {
		str := fmt.Sprintf("%d", i)
		s = append(s, str)
	}

	return s
}

func faster(ints []int) []string {
	s := make([]string, len(ints))
	for ind, i := range ints {
		str := strconv.Itoa(i)
		s[ind] = str
	}

	return s
}
