package commandhandler

import (
	"github.com/szmate1618/go-calc/expressionevaluator"
)

func RunCommand(command string) (string, error) {

	if command == "exit" {
		return "exit", nil
	}

	value, err := expressionevaluator.EvaluateExpression(command)

	return value, err
}
