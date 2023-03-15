package expressionevaluator

func GetParenthesisMask(expression string) string {

	var mask []byte = []byte(expression)

	depth := 0

	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			depth++
		}
		if expression[i] == ')' {
			depth--
			continue
		}
		if depth > 0 {
			mask[i] = '#'
		}
	}

	return string(mask)
}

func IsFullyParenthesized(expression string) bool {

	if expression[0] != '(' {
		return false
	}

	if expression[len(expression)-1] != ')' {
		return false
	}

	depth := 1

	for i := 1; i < len(expression)-1; i++ {
		if expression[i] == '(' {
			depth++
		}
		if expression[i] == ')' {
			depth--
		}
		if depth == 0 {
			return false
		}
	}

	return true
}

func EvaluateExpression(expression string) (string, error) {
	return "", nil
}
