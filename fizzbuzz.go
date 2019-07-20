package main

import (
	"strconv"
)

func Stringify(num int) string {
	var result string

	if num%3 == 0 {
		result += "Fizz"
	}
	if num%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(num)
	}

	return result
}
