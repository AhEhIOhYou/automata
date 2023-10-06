package nfa_test

import (
	"nfa"
	"testing"
)

// Решить задачу об окончаниях -ое, -ая, -ие
func TestNFA(t *testing.T) {
	automata := nfa.NewNFA(3)

	for i := 'А'; i <= 'я'; i++ {
		automata.AddLetter(string(i))
		automata.SetTransition("s0", "s0", string(i))
	}

	automata.SetTransition("s0", "s1", "о")
	automata.SetTransition("s1", "s2", "е")

	automata.SetTransition("s0", "s1", "а")
	automata.SetTransition("s0", "s1", "я")

	automata.SetTransition("s0", "s1", "и")
	automata.SetTransition("s0", "s1", "е")

	automata.SetStartState("s0")
	automata.SetEndState("s2")
}
