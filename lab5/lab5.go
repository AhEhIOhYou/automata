package lab5

import (
	"github.com/ahehiohyou/automata/lab5/etree"
	"github.com/ahehiohyou/automata/lab5/rpn"
	"strconv"
)

func eval(expr *etree.Node, env map[string]float64) float64 {
	switch expr.Token.Type {
	case 0:
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
	case 1:
		return env[expr.Token.Value]
	case 2:
		v, _ := strconv.ParseFloat(expr.Token.Value, 64)
		return v
	default:
		return 0
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
