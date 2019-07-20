package tdd

import "strconv"

func Stringify(num int) string {
	if num%3 == 0 {
		return "fizz"
	}

	str := strconv.Itoa(num)
	return str
}
