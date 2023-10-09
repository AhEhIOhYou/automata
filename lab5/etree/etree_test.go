package etree

import (
	"fmt"
	"testing"
)

func TestCreateTreeFromRPN(t *testing.T) {
	expression := "x1 x2 + y1 * z -"
	//expression := "3 7 + 2 * 10 -"
	tree := CreateTreeFromRPN(expression)
	fmt.Println(tree)
}
