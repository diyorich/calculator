package validators

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateOperands(operands []string) error {
	if !isAllowedOperands(operands) {
		return fmt.Errorf("Check operands, only allowed arabic and rome numerals")
	}

	//проверка что оба операнда римские или арабские
	if !isHomogenous(operands) {
		return fmt.Errorf("Both operands must be or arabic or Roman numerals")
	}

	return nil
}

func isAllowedOperands(operands []string) bool {
	for _, operand := range operands {
		if isRomanNumeral(operand) || isStringInt(operand) {
			continue
		}

		return false
	}

	return true
}

// проверка на однородность
func isHomogenous(operands []string) bool {
	roman := 0
	arabic := 0

	//todo: refactor
	for _, operand := range operands {
		if isRomanNumeral(operand) {
			roman += 1
			continue
		}

		if isStringInt(operand) {
			arabic += 1
			continue
		}
	}

	if !(roman == len(operands) || arabic == len(operands)) {
		return false
	}

	return true
}

func isRomanNumeral(s string) bool {
	var romanNumerals = map[string]int{
		"I": 1, "IV": 4, "V": 5, "IX": 9,
		"X": 10, "XL": 40, "L": 50, "XC": 90,
		"C": 100, "CD": 400, "D": 500, "CM": 900,
		"M": 1000,
	}

	s = strings.ToUpper(s)

	for len(s) > 0 {
		found := false
		for k := range romanNumerals {
			if strings.HasPrefix(s, k) {
				s = s[len(k):]
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func isStringInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
