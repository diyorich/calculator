package handlers

import (
	"calculator/internal/calculator"
	"calculator/internal/validators"
	"calculator/pkg/converter"
	"strconv"
	"strings"
)

const OperatorPosition = 1

var (
	isRoman bool
)

func Handle(operation string) (interface{}, error) {
	//process and initialize operation
	operationElems := strings.Split(operation, " ")
	operator := operationElems[OperatorPosition]
	operationElems = append(operationElems[:OperatorPosition], operationElems[OperatorPosition+1:]...)

	err := validators.ValidateOperator(operator)
	if err != nil {
		return nil, err
	}

	err = validators.ValidateOperands(operationElems)
	if err != nil {
		return nil, err
	}

	var operands []int

	//formatting
	for _, operand := range operationElems {
		if validators.IsRomanNumeral(operand) {
			convertedVal, _ := converter.RomanToInt(operand)
			operands = append(operands, convertedVal)
			isRoman = true
		} else {
			convertedVal, _ := strconv.Atoi(operand)
			operands = append(operands, convertedVal)
			isRoman = false
		}
	}

	//calculation
	answer, _ := calculator.Calculate(operands, operator)

	if isRoman {
		return converter.IntToRoman(answer), nil
	}

	//todo return calculated data

	return answer, nil
}
