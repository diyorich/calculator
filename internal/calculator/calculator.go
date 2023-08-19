package calculator

type Answer interface{}

func Calculate(operands []int, operator string) (Answer, error) {
	var answer Answer

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
