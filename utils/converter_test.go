package utils

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMultipleConvertDecimalToBinary(t *testing.T) {
	errorMessages := ""

	errorMessages += testConvertDecimalToBinary("0", "0")
	errorMessages += testConvertDecimalToBinary("1", "1")
	errorMessages += testConvertDecimalToBinary("2", "10")
	errorMessages += testConvertDecimalToBinary("3", "11")
	errorMessages += testConvertDecimalToBinary("4", "100")
	errorMessages += testConvertDecimalToBinary("5", "101")
	errorMessages += testConvertDecimalToBinary("10", "1010")

	if errorMessages != "" {
		t.Error(errorMessages)
	}
}

func TestMultipleConvertBinaryToDecimal(t *testing.T) {
	errorMessages := ""

	errorMessages += testConvertBinaryToDecimal("0000", "0")
	errorMessages += testConvertBinaryToDecimal("0", "0")
	errorMessages += testConvertBinaryToDecimal("0001", "1")
	errorMessages += testConvertBinaryToDecimal("1", "1")
	errorMessages += testConvertBinaryToDecimal("10", "2")
	errorMessages += testConvertBinaryToDecimal("11", "3")
	errorMessages += testConvertBinaryToDecimal("100", "4")
	errorMessages += testConvertBinaryToDecimal("101", "5")
	errorMessages += testConvertBinaryToDecimal("1010", "10")

	if errorMessages != "" {
		t.Error(errorMessages)
	}
}

func TestMultipleConvertBinaryToDecimalThenConvertDecimalToBinary(t *testing.T) {
	errorMessages := ""

	value := ""
	for i := 0; i < 20; i++ {
		errorMessages += testConvertBinaryToDecimalThenConvertDecimalToBinary(value + strconv.Itoa(i%2))
		if i%3 == 0 {
			value += "1"
		}
	}

	if errorMessages != "" {
		t.Error(errorMessages)
	}
}

func TestMultipleConvertDecimalToBinaryThenConvertBinaryToDecimal(t *testing.T) {
	errorMessages := ""

	for i := 0; i < 100; i++ {
		value := strconv.Itoa(i)
		errorMessages += testConvertDecimalToBinaryThenConvertBinaryToDecimal(value)
	}

	if errorMessages != "" {
		t.Error(errorMessages)
	}
}

func testConvertDecimalToBinary(given string, expected string) string {
	var got = ConvertDecimalToBinary(given)
	if got != expected {
		return fmt.Sprintf("Conversion Decimal to Binary failed!\nGiven:   \t%v\nExpected:\t%v\nGot:     \t%v\n\n", given, expected, got)
	}
	return ""
}

func testConvertBinaryToDecimal(given string, expected string) string {
	var got = ConvertBinaryToDecimal(given)
	if got != expected {
		return fmt.Sprintf("Conversion Binary to Decimal failed!\nGiven:   \t%v\nExpected:\t%v\nGot:     \t%v\n\n", given, expected, got)
	}
	return ""
}

func testConvertBinaryToDecimalThenConvertDecimalToBinary(given string) string {
	var got = ConvertDecimalToBinary(ConvertBinaryToDecimal(given))
	if got != given {
		return fmt.Sprintf("Conversion Binary to Decimal to Binary failed!\nGiven:   \t%v\nExpected:\t%v\nGot:     \t%v\n\n", given, given, got)
	}
	return ""
}

func testConvertDecimalToBinaryThenConvertBinaryToDecimal(given string) string {
	var got = ConvertBinaryToDecimal(ConvertDecimalToBinary(given))
	if got != given {
		return fmt.Sprintf("Conversion Decimal to Binary to Decimal failed!\nGiven:   \t%v\nExpected:\t%v\nGot:     \t%v\n\n", given, given, got)
	}
	return ""
}
