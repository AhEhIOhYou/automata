package pda_test

import (
	"fmt"
	"pda"
	"testing"
)

func TestPDA(t *testing.T) {
	automata := pda.NewPDA(1)

	for i := 'А'; i <= 'я'; i++ {
		automata.AddLetter(string(i))
		automata.SetTransition("s0", "s0", string(i))
	}

	automata.AddLetter("(")
	automata.AddLetter(")")
	automata.SetTransition("s0", "s0", "(")
	automata.SetTransition("s0", "s0", ")")
	automata.AddLetter("{")
	automata.AddLetter("}")
	automata.SetTransition("s0", "s0", "{")
	automata.SetTransition("s0", "s0", "}")
	automata.AddLetter("[")
	automata.AddLetter("]")
	automata.SetTransition("s0", "s0", "[")
	automata.SetTransition("s0", "s0", "]")

	automata.SetStartState("s0")
	automata.SetEndState("s0")

	fmt.Println(automata.Accepts("((({Евреи})))"))
}
