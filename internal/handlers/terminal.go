package handlers

import (
	"calculator/internal/calculator"
	"calculator/internal/validators"
	"calculator/pkg/converter"
	"fmt"
	"os"
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

	if len(operationElems) != 3 {
		fmt.Println("Not enough arguments or operands.")
		os.Exit(1)
	}

	operator := operationElems[OperatorPosition]
	operationElems = append(operationElems[:OperatorPosition], operationElems[OperatorPosition+1:]...)

	err := validators.ValidateOperator(operator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = validators.ValidateOperands(operationElems)
	if err != nil {
		fmt.Println(err)

		//closing program on any error which can occur
		os.Exit(1)
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
		return converter.AnswerToRoman(answer)
	}

	return answer, nil
}
