package lab5_test

import (
	"fmt"
	"github.com/ahehiohyou/automata/lab5"
	"testing"
)

func TestSolve(t *testing.T) {
	expression := "2 + x * ( y + z - y1 ) + w * x2 / w3 - 2"

	env := map[string]float64{
		"x":  3,
		"x2": 7,
		"y":  2,
		"y1": 2,
		"z":  11,
		"w":  1,
		"w3": 5,
	}

	res := lab5.Solve(expression, env)

	if res != 34.4 {
		fmt.Println("Not solved")
	} else {
		fmt.Println("Solved")
	}
}
