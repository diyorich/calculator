package calculator

func Calculate(operands []int, operator string) (int, error) {
	var answer int

	switch operator {
	case "+":
		answer = operands[0] + operands[1]
	case "-":
		answer = operands[0] - operands[1]
	case "*":
		answer = operands[0] * operands[1]
	case "/":
		answer = operands[0] / operands[1]
	}

	return answer, nil
}
