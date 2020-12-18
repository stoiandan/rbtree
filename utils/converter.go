package utils

import (
	"math"
	"strconv"
)

func ConvertDecimalToBinary(decimal string) string {
	number, _ := strconv.Atoi(decimal)
	binary := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = number % 2
		number = number / 2
		binary += remainder * counter
		counter *= 10

	}
	return strconv.Itoa(binary)
}

func ConvertBinaryToDecimal(binary string) string {
	number, _ := strconv.Atoi(binary)
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return strconv.Itoa(decimal)
}
