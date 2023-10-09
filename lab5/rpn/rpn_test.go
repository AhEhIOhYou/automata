package rpn

import (
	"fmt"
	"testing"
)

func TestToRPN(t *testing.T) {
	//expression := "( 3 + 7 ) * 2 - 10 / 2"
	expression := "( x1 + x2 ) * y1 - y2 / z1"
	result := ToRPN(expression)
	fmt.Println("RPN:", result)
}
