package validators

import (
	"calculator/pkg/converter"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateOperands(operands []string) error {
	if len(operands) < 0 {
		return fmt.Errorf("amount of operands must be more than 1")
	}

	if len(operands) == 1 {
		return fmt.Errorf("amount of operands must be more than 1")
	}

	if len(operands) > 2 {
		return fmt.Errorf("amount of operands must be 2")
	}

	if !isAllowedOperands(operands) {
		return fmt.Errorf("check operands, only allowed pair of Arabic or pair of Roman numerals")
	}

	//проверка что оба операнда римские или арабские
	if !isHomogenous(operands) {
		fmt.Println("operands must be pair of Romans or pair of Numerals")
		os.Exit(1)
	}

	if !isAllowedRangeOfNumber(operands) {
		return fmt.Errorf("calculator accepts number in range [1:10]")
	}

	return nil
}

func isAllowedRangeOfNumber(operands []string) bool {
	var formatted int
	for _, operand := range operands {
		if IsRomanNumeral(operand) {
			formatted, _ = converter.RomanToInt(operand)
		} else {
			formatted, _ = strconv.Atoi(operand)
		}

		if formatted > 10 || formatted < 1 {
			return false
		}
	}

	return true
}

func isAllowedOperands(operands []string) bool {
	for _, operand := range operands {
		if IsRomanNumeral(operand) || isStringInt(operand) {
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
		if IsRomanNumeral(operand) {
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

func IsRomanNumeral(s string) bool {
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
