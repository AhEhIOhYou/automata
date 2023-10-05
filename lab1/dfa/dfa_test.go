package dfa_test

import (
	"dfa"
	"fmt"
	"testing"
)

func TestDfa(t *testing.T) {

}

func TestLab3(t *testing.T) {
	fmt.Println(dfa.EmailCheck("vladimirov_d1ma@mail.ru"))

	fmt.Println(dfa.EmailCheck("xd"))
}
