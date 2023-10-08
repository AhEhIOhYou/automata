package pda

import (
	"container/list"
	"fmt"
)

// State представляет состояние КАМП
type State struct {
	name string
	term bool
}

// NewState создает новое состояние с заданным именем и флагом заключительности
func NewState(name string, term bool) *State {
	return &State{name: name, term: term}
}

// IsTerminal возвращает true, если состояние является заключительным
func (s *State) IsTerminal() bool {
	return s.term
}

type Letter struct {
	name string
}

func NewLetter(name string) *Letter {
	return &Letter{name: name}
}

type PDA struct {
	states  map[*State]bool
	letters map[*Letter]bool
	trans   map[*State]map[*Letter]*State
	start   *State
	current *State
	stack   *list.List
}

func NewPDA(statesCount int) *PDA {
	pda := PDA{
		states:  make(map[*State]bool),
		letters: make(map[*Letter]bool),
		trans:   make(map[*State]map[*Letter]*State),
		stack:   list.New(),
	}

	if statesCount < 0 {
		return nil
	} else if statesCount == 0 {
		return &pda
	}

	for i := 0; i < statesCount; i++ {
		name := fmt.Sprintf("s%d", i)
		pda.AddState(name, false)
	}

	return &pda
}

// AddState добавляет новое состояние в КАМП с заданным именем и флагом заключительности
// Возвращает указатель на добавленное состояние или nil, если такое имя уже существует
func (p *PDA) AddState(name string, term bool) *State {
	for s := range p.states {
		if s.name == name {
			return nil // имя уже занято
		}
	}
	state := NewState(name, term)
	p.states[state] = true
	p.trans[state] = make(map[*Letter]*State)
	return state
}

// RemoveState удаляет заданное состояние из КАМП и все связанные с ним переходы
// Возвращает true, если удаление прошло успешно, или false, если такого состояния не существует
func (p *PDA) RemoveState(state *State) bool {
	if _, ok := p.states[state]; !ok {
		return false // такого состояния нет в ДКА
	}
	delete(p.states, state)
	delete(p.trans, state)
	for _, m := range p.trans {
		for l := range m {
			if m[l] == state {
				delete(m, l)
			}
		}
	}
	if p.start == state {
		p.start = nil
	}
	if p.current == state {
		p.current = nil
	}
	return true
}

// AddLetter добавляет новый символ в алфавит КАМП с заданным именем
// Возвращает указатель на добавленный символ или nil, если такое имя уже существует
func (p *PDA) AddLetter(name string) *Letter {
	for l := range p.letters {
		if l.name == name {
			return nil // имя уже занято
		}
	}
	letter := NewLetter(name)
	p.letters[letter] = true
	return letter
}

// RemoveLetter удаляет заданный символ из алфавита КАМП и все связанные с ним переходы
// Возвращает true, если удаление прошло успешно, или false, если такого символа не существует
func (p *PDA) RemoveLetter(letter *Letter) bool {
	if _, ok := p.letters[letter]; !ok {
		return false // такого символа нет в алфавите ДКА
	}
	delete(p.letters, letter)
	for _, m := range p.trans {
		delete(m, letter) // удалить переход по удаляемому символу
	}
	return true
}

// FindLetterByName возвращает ссылку на букву алфавита по её имени
func (p *PDA) FindLetterByName(name string) *Letter {
	for letter := range p.letters {
		if letter.name == name {
			return letter
		}
	}
	return nil
}

// FindStateByName возвращает ссылку на состояние по имени.
// Возвращает nil если состояние не принадлежит КАМП и ссылку на состояние если принадлежит
func (p *PDA) FindStateByName(name string) *State {
	for state := range p.states {
		if state.name == name {
			return state
		}
	}
	return nil
}

// SetStartState устанавливает начальное состояние КАМП
// Возвращает true, если состояние установлено успешно, или false, если заданное состояние не принадлежит КАМП
func (p *PDA) SetStartState(name string) bool {
	s := p.FindStateByName(name)
	if s == nil {
		return false
	}
	p.start = s
	p.current = s
	return true
}

// SetEndState устанавливает состояние как конечное
func (p *PDA) SetEndState(name string) bool {
	s := p.FindStateByName(name)
	if s == nil {
		return false
	}
	s.term = true
	return true
}

// IsEndState проверка, является ли состояние конечным
func (p *PDA) IsEndState() bool {
	return p.current.IsTerminal()
}

// GetCurrentState возвращает текущее состояние КАМП или nil, если оно не установлено
func (p *PDA) GetCurrentState() *State {
	return p.current
}

// ResetCurrentState сбрасывает текущее состояние КАМП в начальное состояние
func (p *PDA) ResetCurrentState() {
	p.stack = list.New()
	p.current = p.start
}

// PushStack добавляет в конец стека новый элемент
func (p *PDA) PushStack(element string) {
	p.stack.PushBack(element)
}

// PopStack возвращает последний элемент стека
func (p *PDA) PopStack() string {
	e := p.stack.Back()
	if e != nil {
		p.stack.Remove(e)
		if str, ok := e.Value.(string); ok {
			return str
		}
	}
	return ""
}

// IsStackEmpty проверяет пустой ли сейчас стек
func (p *PDA) IsStackEmpty() bool {
	return p.stack.Len() == 0
}

// SetTransition устанавливает переход из заданного исходного состояния в заданное конечное состояние по заданному символу
// Возвращает true, если переход установлен успешно, или false, если какой-то из параметров не принадлежит КАМП
func (p *PDA) SetTransition(fromName, toName, letterBy string) bool {
	by := p.FindLetterByName(letterBy)
	from := p.FindStateByName(fromName)
	to := p.FindStateByName(toName)
	if _, ok := p.states[from]; !ok {
		return false // исходное состояние не принадлежит ДКА
	}
	if _, ok := p.states[to]; !ok {
		return false // конечное состояние не принадлежит ДКА
	}
	if _, ok := p.letters[by]; !ok {
		return false // символ не принадлежит алфавиту ДКА
	}
	p.trans[from][by] = to
	return true
}

// RemoveTransition удаляет переход из заданного исходного состояния по заданному символу
// Возвращает true, если переход удален успешно, или false, если какой-то из параметров не принадлежит КАМП или перехода не существует
func (p *PDA) RemoveTransition(from *State, by *Letter) bool {
	if _, ok := p.states[from]; !ok {
		return false // исходное состояние не принадлежит ДКА
	}
	if _, ok := p.letters[by]; !ok {
		return false // символ не принадлежит алфавиту ДКА
	}
	if _, ok := p.trans[from][by]; !ok {
		return false // перехода не существует
	}
	delete(p.trans[from], by) // удалить переход
	return true
}

// Transition выполняет переход из текущего состояния в другое по заданному символу и возвращает новое текущее состояние
func (p *PDA) Transition(by *Letter) *State {
	if p.current == nil {
		return nil // текущее состояние не установлено
	}
	if _, ok := p.letters[by]; !ok {
		return nil // символ не принадлежит алфавиту ДКА
	}
	if to, ok := p.trans[p.current][by]; ok {
		p.current = to
		return p.current
	}
	return nil
}

// Accepts проверяет строку на принадлежность языку КАМП
// Возвращает true, если строка принадлежит языку КАМП, или false, если нет
func (p *PDA) Accepts(s string) bool {

	braces := []struct {
		open  string
		close string
	}{
		{
			open:  "(",
			close: ")",
		},
		{
			open:  "{",
			close: "}",
		},
		{
			open:  "[",
			close: "]",
		},
	}

	p.ResetCurrentState()

	for _, r := range s {
		r := string(r)

		l := p.FindLetterByName(r)
		if l == nil {
			return false
		}

		for _, b := range braces {
			if b.open == r {
				p.PushStack(b.close)
			} else if b.close == r {
				if p.PopStack() != r {
					return false
				}
			}
		}

		if p.Transition(l) == nil {
			return false
		}
	}

	return p.IsEndState() && p.IsStackEmpty()
}
