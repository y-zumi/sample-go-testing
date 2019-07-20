package tdd

import "strconv"

func Stringify(num int) string {
	if num%3 == 0 {
		return "fizz"
	}

	if num%5 == 0 {
		return "buzz"
	}

	str := strconv.Itoa(num)
	return str
}
