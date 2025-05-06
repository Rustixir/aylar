package tool

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c *Calculator) Name() string {
	return "Calculator"
}

func (c *Calculator) Description() string {
	return "Evaluates simple math expressions like '2 + 2'."
}

func (c *Calculator) Run(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid expression format, use 'a + b'")
	}

	a, err1 := strconv.ParseFloat(parts[0], 64)
	op := parts[1]
	b, err2 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil {
		return "", fmt.Errorf("invalid numbers")
	}

	switch op {
	case "+":
		return fmt.Sprintf("%f", a+b), nil
	case "-":
		return fmt.Sprintf("%f", a-b), nil
	case "*":
		return fmt.Sprintf("%f", a*b), nil
	case "/":
		if b == 0 {
			return "", fmt.Errorf("division by zero")
		}
		return fmt.Sprintf("%f", a/b), nil
	default:
		return "", fmt.Errorf("unknown operator")
	}
}
