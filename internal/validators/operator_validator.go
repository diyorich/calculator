package validators

import (
	"fmt"
)

func ValidateOperator(operator string) error {
	if !operatorExist(operator) {
		return fmt.Errorf("operation %s does not exist", operator)
	}

	return nil
}

func operatorExist(operator string) bool {
	allowedOperators := [4]string{"+", "-", "*", "/"}
	found := false
	for _, op := range allowedOperators {
		if op == operator {
			found = true
		}
	}

	return found
}
