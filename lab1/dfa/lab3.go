package dfa

import "strconv"

func enAlpAdd(aut *DFA) {
	for i := 97; i < 123; i++ {
		aut.AddLetter(string(rune(i)))
	}
}

func numsAlpAdd(aut *DFA) {
	for i := 0; i < 10; i++ {
		aut.AddLetter(strconv.Itoa(i))
	}
}

func dotAplAdd(aut *DFA) {
	aut.AddLetter(".")
}

func atAplAdd(aut *DFA) {
	aut.AddLetter("@")
}

func symAplAdd(aut *DFA) {
	aut.AddLetter("_")
	aut.AddLetter("-")
}

func transitionsAdd(aut *DFA) {
	for i := 97; i < 123; i++ {
		aut.SetTransition("s0", "s1", string(rune(i)))
		aut.SetTransition("s1", "s1", string(rune(i)))
		aut.SetTransition("s2", "s2", string(rune(i)))
	}

	for i := 0; i < 10; i++ {
		aut.SetTransition("s1", "s1", strconv.Itoa(i))
		aut.SetTransition("s2", "s2", strconv.Itoa(i))
	}

	aut.SetTransition("s1", "s1", "_")
	aut.SetTransition("s2", "s2", "_")

	aut.SetTransition("s1", "s1", "-")
	aut.SetTransition("s2", "s2", "-")

	aut.SetTransition("s1", "s2", "@")
	aut.SetTransition("s2", "s3", ".")

	aut.SetTransition("s3", "s4", "c")
	aut.SetTransition("s4", "s5", "o")
	aut.SetTransition("s5", "s6", "m")

	aut.SetTransition("s3", "s7", "r")
	aut.SetTransition("s7", "s8", "u")
}

func EmailCheck(s string) bool {
	if s == "" {
		return false
	}

	automata := NewDFA(9)

	enAlpAdd(automata)
	numsAlpAdd(automata)
	dotAplAdd(automata)
	atAplAdd(automata)
	symAplAdd(automata)

	transitionsAdd(automata)

	automata.SetStartState("s0")
	automata.SetEndState("s6")
	automata.SetEndState("s8")

	return automata.Accepts(s)
}
