package dfa

import "fmt"

// State представляет состояние ДКА
type State struct {
	name string
	term bool
}

// NewState создает новое состояние с заданным именем и флагом заключительности
func NewState(name string, term bool) *State {
	return &State{name: name, term: term}
}

// String возвращает строковое представление состояния
func (s *State) String() string {
	return s.name
}

// IsTerminal возвращает true, если состояние является заключительным
func (s *State) IsTerminal() bool {
	return s.term
}

// Letter представляет символ алфавита ДКА
type Letter struct {
	name string // имя символа
}

// NewLetter создает новый символ с заданным именем
func NewLetter(name string) *Letter {
	return &Letter{name: name}
}

// String возвращает строковое представление символа
func (l *Letter) String() string {
	return l.name
}

// DFA представляет детерминированный конечный автомат
type DFA struct {
	states  map[*State]bool               // множество состояний ДКА
	letters map[*Letter]bool              // множество символов алфавита ДКА
	trans   map[*State]map[*Letter]*State // функция переходов ДКА
	start   *State                        // начальное состояние ДКА
	current *State                        // текущее состояние ДКА
}

// NewDFA создает новый ДКА
func NewDFA(statesCount int) *DFA {
	dfa := DFA{
		states:  make(map[*State]bool),
		letters: make(map[*Letter]bool),
		trans:   make(map[*State]map[*Letter]*State),
	}

	if statesCount < 0 {
		return nil
	} else if statesCount == 0 {
		return &dfa
	}

	for i := 0; i < statesCount; i++ {
		name := fmt.Sprintf("s%d", i)
		dfa.AddState(name, false)
	}

	return &dfa
}

// AddState добавляет новое состояние в ДКА с заданным именем и флагом заключительности
// Возвращает указатель на добавленное состояние или nil, если такое имя уже существует
func (d *DFA) AddState(name string, term bool) *State {
	for s := range d.states {
		if s.name == name {
			return nil // имя уже занято
		}
	}
	state := NewState(name, term)
	d.states[state] = true
	d.trans[state] = make(map[*Letter]*State)
	return state
}

// RemoveState удаляет заданное состояние из ДКА и все связанные с ним переходы
// Возвращает true, если удаление прошло успешно, или false, если такого состояния не существует
func (d *DFA) RemoveState(state *State) bool {
	if _, ok := d.states[state]; !ok {
		return false // такого состояния нет в ДКА
	}
	delete(d.states, state)
	delete(d.trans, state)
	for _, m := range d.trans {
		for l := range m {
			if m[l] == state {
				delete(m, l) // удалить переход в удаляемое состояние
			}
		}
	}
	if d.start == state {
		d.start = nil // обнулить начальное состояние, если оно удаляется
	}
	if d.current == state {
		d.current = nil // обнулить текущее состояние, если оно удаляется
	}
	return true
}

// AddLetter добавляет новый символ в алфавит ДКА с заданным именем
// Возвращает указатель на добавленный символ или nil, если такое имя уже существует
func (d *DFA) AddLetter(name string) *Letter {
	for l := range d.letters {
		if l.name == name {
			return nil // имя уже занято
		}
	}
	letter := NewLetter(name)
	d.letters[letter] = true
	return letter
}

// RemoveLetter удаляет заданный символ из алфавита ДКА и все связанные с ним переходы
// Возвращает true, если удаление прошло успешно, или false, если такого символа не существует
func (d *DFA) RemoveLetter(letter *Letter) bool {
	if _, ok := d.letters[letter]; !ok {
		return false // такого символа нет в алфавите ДКА
	}
	delete(d.letters, letter)
	for _, m := range d.trans {
		delete(m, letter) // удалить переход по удаляемому символу
	}
	return true
}

// SetTransition устанавливает переход из заданного исходного состояния в заданное конечное состояние по заданному символу
// Возвращает true, если переход установлен успешно, или false, если какой-то из параметров не принадлежит ДКА
func (d *DFA) SetTransition(fromName, toName, letterBy string) bool {
	by := d.FindLetterByName(letterBy)
	from := d.FindStateByName(fromName)
	to := d.FindStateByName(toName)
	if _, ok := d.states[from]; !ok {
		return false // исходное состояние не принадлежит ДКА
	}
	if _, ok := d.states[to]; !ok {
		return false // конечное состояние не принадлежит ДКА
	}
	if _, ok := d.letters[by]; !ok {
		return false // символ не принадлежит алфавиту ДКА
	}
	d.trans[from][by] = to // установить переход
	return true
}

// FindLetterByName возвращает ссылку на букву алфавита по её имени
func (d *DFA) FindLetterByName(name string) *Letter {
	for letter := range d.letters {
		if letter.name == name {
			return letter
		}
	}
	return nil
}

// FindStateByName возвращает ссылку на состояние по имени.
// Возвращает nil если состояние не принадлежит ДКА и ссылку на состояние если принадлежит
func (d *DFA) FindStateByName(name string) *State {
	for state := range d.states {
		if state.name == name {
			return state
		}
	}
	return nil
}

// RemoveTransition удаляет переход из заданного исходного состояния по заданному символу
// Возвращает true, если переход удален успешно, или false, если какой-то из параметров не принадлежит ДКА или перехода не существует
func (d *DFA) RemoveTransition(from *State, by *Letter) bool {
	if _, ok := d.states[from]; !ok {
		return false // исходное состояние не принадлежит ДКА
	}
	if _, ok := d.letters[by]; !ok {
		return false // символ не принадлежит алфавиту ДКА
	}
	if _, ok := d.trans[from][by]; !ok {
		return false // перехода не существует
	}
	delete(d.trans[from], by) // удалить переход
	return true
}

// SetStartState устанавливает начальное состояние ДКА
// Возвращает true, если состояние установлено успешно, или false, если заданное состояние не принадлежит ДКА
func (d *DFA) SetStartState(name string) bool {
	state := d.FindStateByName(name)
	if state == nil {
		return false
	}
	d.start = state   // установить начальное состояние
	d.current = state // установить текущее состояние равным начальному
	return true
}

// SetEndState устанавливает состояние как конечное
func (d *DFA) SetEndState(name string) bool {
	s := d.FindStateByName(name)
	if s == nil {
		return false
	}
	s.term = true
	return true
}

// GetStartState возвращает начальное состояние ДКА или nil, если оно не установлено
func (d *DFA) GetStartState() *State {
	return d.start
}

// IsEndState возвращает true если текущее состояние автомата заключительное
func (d *DFA) IsEndState() bool {
	return d.current.IsTerminal()
}

// GetCurrentState возвращает текущее состояние ДКА или nil, если оно не установлено
func (d *DFA) GetCurrentState() *State {
	return d.current
}

// ResetCurrentState сбрасывает текущее состояние ДКА в начальное состояние или nil, если оно не установлено
func (d *DFA) ResetCurrentState() {
	d.current = d.start
}

// Transition выполняет переход из текущего состояния в другое по заданному символу и возвращает новое текущее состояние
func (d *DFA) Transition(by *Letter) *State {
	if d.current == nil {
		return nil // текущее состояние не установлено
	}
	if _, ok := d.letters[by]; !ok {
		return nil // символ не принадлежит алфавиту ДКА
	}
	if to, ok := d.trans[d.current][by]; ok {
		d.current = to // выполнить переход
		return d.current
	}
	return nil // перехода не существует
}

// CheckChain проверяет цепочку символов на принадлежность языку ДКА
// Возвращает true, если цепочка принадлежит языку ДКА, или false, если нет
func (d *DFA) CheckChain(chain []*Letter) bool {
	d.ResetCurrentState()
	for _, l := range chain {
		if d.Transition(l) == nil {
			return false
		}
	}
	return d.IsEndState()
}

// Accepts проверяет строку на принадлежность языку ДКА
// Возвращает true, если строка принадлежит языку ДКА, или false, если нет
func (d *DFA) Accepts(s string) bool {
	d.ResetCurrentState()
	for _, r := range s {
		var l *Letter
		l = d.FindLetterByName(string(r))
		if l == nil {
			return false
		}
		if d.Transition(l) == nil {
			return false
		}
	}
	return d.IsEndState()
}
