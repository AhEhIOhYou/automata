package dfa_test

import (
	"dfa"
	"fmt"
	"testing"
)

func TestDfa(t *testing.T) {

	automata := dfa.NewDFA()

	a := automata.AddLetter("0")
	b := automata.AddLetter("1")
	c := automata.AddLetter("2")
	d := automata.AddLetter("3")
	e := automata.AddLetter("4")
	f := automata.AddLetter("5")
	g := automata.AddLetter("6")
	h := automata.AddLetter("7")
	i := automata.AddLetter("8")

	s0 := automata.AddState("s0", false)
	s1 := automata.AddState("s1", false)
	s2 := automata.AddState("s2", false)
	s3 := automata.AddState("s3", false)
	s4 := automata.AddState("s4", false)
	s5 := automata.AddState("s5", true)
	s6 := automata.AddState("s6", true)

	automata.SetTransition(s0, s3, a)
	automata.SetTransition(s0, s1, b)
	automata.SetTransition(s0, s6, h)
	automata.SetTransition(s1, s3, d)
	automata.SetTransition(s1, s2, c)
	automata.SetTransition(s3, s2, a)
	automata.SetTransition(s3, s5, f)
	automata.SetTransition(s2, s4, e)
	automata.SetTransition(s2, s2, i)
	automata.SetTransition(s4, s5, g)
	automata.SetTransition(s4, s6, h)
	automata.SetTransition(s6, s2, d)

	automata.SetStartState(s0)

	states := []*dfa.Letter{b, d, f}

	fmt.Println(automata.CheckChain(states))

	fmt.Println(automata.Accepts("a135"))
}
