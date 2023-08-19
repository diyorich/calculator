package converter

import (
	"fmt"
	"strings"
)

var (
	romanSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	romanToInteger = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func RomanToInt(roman string) (int, error) {
	roman = strings.ToUpper(roman)
	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanToInteger[string(roman[i])]
		if value == 0 {
			return 0, fmt.Errorf("invalid Roman numeral character: %s", string(roman[i]))
		}

		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result, nil
}

func IntToRoman(num int) string {
	var result strings.Builder

	for _, symbol := range romanSymbols {
		for num >= symbol.value {
			result.WriteString(symbol.symbol)
			num -= symbol.value
		}
	}

	return result.String()
}
