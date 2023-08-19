package handlers

import (
	"calculator/internal/validators"
	"strings"
)

const OperatorPosition = 1

func Handle(operation string) error {
	//process and initialize operation
	operationElems := strings.Split(operation, " ")
	operator := operationElems[OperatorPosition]
	operationElems = append(operationElems[:OperatorPosition], operationElems[OperatorPosition+1:]...)

	err := validators.ValidateOperator(operator)
	if err != nil {
		return err
	}

	err = validators.ValidateOperands(operationElems)
	if err != nil {
		return err
	}

	return nil
}
