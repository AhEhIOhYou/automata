package lab5

import (
	"github.com/ahehiohyou/automata/lab5/etree"
	"github.com/ahehiohyou/automata/lab5/rpn"
)

func eval(expr *etree.Node, env map[string]float64) float64 {
	switch expr.Token.Type {
	case 0:
		return env[expr.Token.Value]
	case 1:
		switch expr.Token.Value {
		case "+":
			return eval(expr.Left, env) + eval(expr.Right, env)
		case "-":
			return eval(expr.Left, env) - eval(expr.Right, env)
		case "*":
			return eval(expr.Left, env) * eval(expr.Right, env)
		case "/":
			return eval(expr.Left, env) / eval(expr.Right, env)
		}
	}

	return 0
}

func Solve(expression string, env map[string]float64) float64 {
	res := 0.0

	f := rpn.ToRPN(expression)
	t := etree.CreateTreeFromRPN(f)

	res = eval(t, env)

	return res
}
