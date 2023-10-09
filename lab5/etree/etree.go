// Package etree - для составления бинарного дерева выражения из ОПФ
package etree

import (
	"strings"
)

type TokenType int

const (
	Variable TokenType = iota
	Operator
)

type Token struct {
	Type  TokenType
	Value string
}

type Node struct {
	Token Token
	Left  *Node
	Right *Node
}

func NewNode(t Token) *Node {
	return &Node{Token: t}
}

func IsOperator(value string) bool {
	return strings.ContainsAny(value, "+-*/")
}

func CreateTreeFromRPN(expression string) *Node {
	stack := []*Node{}

	for _, token := range strings.Fields(expression) {
		t := Token{Value: token}
		if IsOperator(token) {
			t.Type = Operator
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Node{Token: t, Left: left, Right: right})
		} else {
			t.Type = Variable
			stack = append(stack, NewNode(t))
		}
	}

	return stack[0]
}
