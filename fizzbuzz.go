package tdd

import "strconv"

func Stringify(num int) string {
	var result string

	if num%3 == 0 {
		result += "fizz"
	}
	if num%5 == 0 {
		result += "buzz"
	}
	if result == "" {
		result = strconv.Itoa(num)
	}

	return result
}
