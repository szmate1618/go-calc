package expression

func GetParenthesisMask(expression string) string {

	var mask []byte = []byte(expression)

	depth := 0

	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			depth++
		}
		if expression[i] == ')' {
			depth--
		}
		if depth > 0 && expression[i] != ')' {
			mask[i] = '#'
		}
	}

	return string(mask)
}

func IsFullyParenthesized(expression string) bool {

	if expression == "" {
		return false
	}

	mask := GetParenthesisMask(expression)

	for i := 0; i < len(mask); i++ {
		if mask[i] != '#' {
			return false
		}
	}

	return true
}

func ParenthesisStripped(expression string) string {
	for IsFullyParenthesized(expression) == false {
		expression = expression[1 : len(expression)-1]
	}
	return expression
}

func EvaluateExpression(expression string) (string, error) {
	return "", nil
}
