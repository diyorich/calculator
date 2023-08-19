package validators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name           string
	Value          []string
	ExpectedError  string
	ExpectedResult interface{}
}

func TestOperandsNegativeCases(t *testing.T) {
	testCases := []TestCase{
		{
			"test both invalid operands",
			[]string{"S", "SOOOOOS"},
			"check operands, only allowed pair of Arabic or pair of Roman numerals",
			nil,
		},
		{
			"test first roman second numeral input",
			[]string{"I", "1"},
			"operands must be pair of Romans or pair of Numerals",
			nil,
		},
		{
			"test first numeral second roman input",
			[]string{"1", "II"},
			"operands must be pair of Romans or pair of Numerals",
			nil,
		},
		{
			"test first numeral max value is more than allowed",
			[]string{"11", "1"},
			"calculator accepts number in range [1:10]",
			nil,
		},
		{
			"test second numeral max value is more than allowed",
			[]string{"1", "1123"},
			"calculator accepts number in range [1:10]",
			nil,
		},
		{
			"test first numeral negative value is less than allowed",
			[]string{"-1", "1"},
			"calculator accepts number in range [1:10]",
			nil,
		},
		{
			"test second numeral max value is more than allowed",
			[]string{"1", "-1001"},
			"calculator accepts number in range [1:10]",
			nil,
		},
		{
			"first roman value is more than allowed",
			[]string{"XX", "I"},
			"calculator accepts number in range [1:10]",
			nil,
		},
		{
			"second roman value is more than allowed",
			[]string{"X", "XXV"},
			"calculator accepts number in range [1:10]",
			nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := ValidateOperands(testCase.Value)
			expectedError := errors.New(testCase.ExpectedError)

			assert.Equal(t, err, expectedError)
		})
	}
}
