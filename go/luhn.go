package main

import (
	"strings"
)

func Valid(number string) bool {
	// strip spaces
	number = strings.Replace(number, " ", "", -1)

	// check if the number is too short
	if len(number) <= 1 {
		return false
	}

	var sum int
	toDouble := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')

		if digit < 0 || digit > 9 {
			return false
		}

		if toDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		toDouble = !toDouble
		sum += digit
	}

	return sum%10 == 0
}

func main() {
	println(Valid("059"))
}
