package nfa

import (
	"fmt"
	"testing"
)

func TestNFA(t *testing.T) {
	automata := NewNFA()
	s0 := automata.AddState("s0", false)
	s1 := automata.AddState("s1", false)
	s2 := automata.AddState("s2", false)
	s3 := automata.AddState("s3", true)

	chern := automata.AddLetter("черн")
	a := automata.AddLetter("а")
	e := automata.AddLetter("е")
	i := automata.AddLetter("и")
	o := automata.AddLetter("о")
	ya := automata.AddLetter("я")

	// Добавить переходы распознающий окончания -ое, -ая, -ие
	automata.AddTransition(s0, s1, chern) // переход из s0 в s1 по символу черн

	automata.AddTransition(s1, s2, a)
	automata.AddTransition(s1, s2, o)
	automata.AddTransition(s1, s2, i)

	automata.AddTransition(s2, s3, e)
	automata.AddTransition(s2, s3, ya)

	automata.SetStartState(s0)

	states := []*Letter{chern, o, e}
	fmt.Println(automata.RecognizeArray(states))

	states = []*Letter{chern, o, i, ya}
	fmt.Println(automata.RecognizeArray(states))
}
