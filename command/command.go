package command

import (
	"github.com/szmate1618/go-calc/expression"
)

func RunCommand(command string) (string, error) {

	if command == "exit" {
		return "exit", nil
	}

	value, err := expression.EvaluateExpression(command)

	return value, err
}
