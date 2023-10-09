package rpn

import (
	"strings"
)

type Operator struct {
	Symbol     string
	Precedence int
}

var Operators = map[string]Operator{
	"+": Operator{Symbol: "+", Precedence: 1},
	"-": Operator{Symbol: "-", Precedence: 1},
	"*": Operator{Symbol: "*", Precedence: 2},
	"/": Operator{Symbol: "/", Precedence: 2},
}

func ToRPN(expression string) string {
	stack := []string{}
	output := []string{}

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if _, ok := Operators[token]; ok {
			for len(stack) > 0 && Operators[stack[len(stack)-1]].Precedence >= Operators[token].Precedence && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else {
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return strings.Join(output, " ")
}
