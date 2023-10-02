package nfa

// State представляет состояние НКА
type State struct {
	name string
	term bool // является ли состояние заключительным
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

// Letter представляет символ алфавита НКА
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

// NFA представляет недетерминированный конечный автомат
type NFA struct {
	states  map[*State]bool                 // множество состояний НКА
	letters map[*Letter]bool                // множество символов алфавита НКА
	trans   map[*State]map[*Letter][]*State // функция переходов НКА
	start   *State                          // начальное состояние НКА
	current []*State                        // текущее множество состояний НКА
}

// NewNFA создает новый НКА без состояний и переходов
func NewNFA() *NFA {
	return &NFA{
		states:  make(map[*State]bool),
		letters: make(map[*Letter]bool),
		trans:   make(map[*State]map[*Letter][]*State),
	}
}

// AddState добавляет новое состояние в НКА с заданным именем и флагом заключительности
// Возвращает указатель на добавленное состояние или nil, если такое имя уже существует
func (n *NFA) AddState(name string, term bool) *State {
	for s := range n.states {
		if s.name == name {
			return nil // имя уже занято
		}
	}
	state := NewState(name, term)
	n.states[state] = true
	n.trans[state] = make(map[*Letter][]*State)
	return state
}

// RemoveState удаляет заданное состояние из НКА и все связанные с ним переходы
// Возвращает true, если удаление прошло успешно, или false, если такого состояния не существует
func (n *NFA) RemoveState(state *State) bool {
	if _, ok := n.states[state]; !ok {
		return false // такого состояния нет в НКА
	}
	delete(n.states, state)
	delete(n.trans, state)
	for _, m := range n.trans {
		for l := range m {
			for i := 0; i < len(m[l]); i++ {
				if m[l][i] == state {
					m[l] = append(m[l][:i], m[l][i+1:]...) // удалить переход в удаляемое состояние
					i--
				}
			}
		}
	}
	if n.start == state {
		n.start = nil
	}
	for i := 0; i < len(n.current); i++ {
		if n.current[i] == state {
			n.current = append(n.current[:i], n.current[i+1:]...) // удалить состояние из текущего множества, если оно удаляется
			i--
		}
	}
	return true
}

// AddLetter добавляет новый символ в алфавит НКА с заданным именем
// Возвращает указатель на добавленный символ или nil, если такое имя уже существует
func (n *NFA) AddLetter(name string) *Letter {
	for l := range n.letters {
		if l.name == name {
			return nil // имя уже занято
		}
	}
	letter := NewLetter(name)
	n.letters[letter] = true
	return letter
}

// RemoveLetter удаляет заданный символ из алфавита НКА и все связанные с ним переходы
// Возвращает true, если удаление прошло успешно, или false, если такого символа не существует
func (n *NFA) RemoveLetter(letter *Letter) bool {
	if _, ok := n.letters[letter]; !ok {
		return false // такого символа нет в алфавите НКА
	}
	delete(n.letters, letter)
	for _, m := range n.trans {
		delete(m, letter)
	}
	return true
}

// AddTransition добавляет переход из заданного исходного состояния в заданное конечное состояние по заданному символу
// Возвращает true, если переход добавлен успешно, или false, если какой-то из параметров не принадлежит НКА
func (n *NFA) AddTransition(from, to *State, by *Letter) bool {
	if _, ok := n.states[from]; !ok {
		return false // исходное состояние не принадлежит НКА
	}
	if _, ok := n.states[to]; !ok {
		return false // конечное состояние не принадлежит НКА
	}
	if _, ok := n.letters[by]; !ok {
		return false // символ не принадлежит алфавиту НКА
	}
	n.trans[from][by] = append(n.trans[from][by], to)
	return true
}

// RemoveTransition удаляет переход из заданного исходного состояния в заданное конечное состояние по заданному символу
// Возвращает true, если переход удален успешно, или false, если какой-то из параметров не принадлежит НКА или перехода не существует
func (n *NFA) RemoveTransition(from, to *State, by *Letter) bool {
	if _, ok := n.states[from]; !ok {
		return false // исходное состояние не принадлежит НКА
	}
	if _, ok := n.states[to]; !ok {
		return false // конечное состояние не принадлежит НКА
	}
	if _, ok := n.letters[by]; !ok {
		return false // символ не принадлежит алфавиту НКА
	}
	for i := 0; i < len(n.trans[from][by]); i++ {
		if n.trans[from][by][i] == to {
			n.trans[from][by] = append(n.trans[from][by][:i], n.trans[from][by][i+1:]...) // удалить переход
			i--
			return true
		}
	}
	return false // перехода не существует
}

// SetStartState устанавливает начальное состояние НКА
// Возвращает true, если состояние установлено успешно, или false, если заданное состояние не принадлежит НКА
func (n *NFA) SetStartState(state *State) bool {
	if _, ok := n.states[state]; !ok {
		return false
	}
	n.start = state
	n.current = []*State{state}
	return true
}

// GetStartState возвращает начальное состояние НКА или nil, если оно не установлено
func (n *NFA) GetStartState() *State {
	return n.start
}

// GetCurrentStates возвращает текущее множество состояний НКА или nil, если оно не установлено
func (n *NFA) GetCurrentStates() []*State {
	return n.current
}

// ResetCurrentStates сбрасывает текущее множество состояний НКА в начальное состояние или nil, если оно не установлено
func (n *NFA) ResetCurrentStates() {
	n.current = []*State{n.start}
}

// Transition выполняет переход из текущего множества состояний в другое по заданному символу и возвращает новое текущее множество состояний
func (n *NFA) Transition(by *Letter) []*State {
	if len(n.current) == 0 {
		return nil
	}
	if _, ok := n.letters[by]; !ok {
		return nil
	}
	next := make(map[*State]bool)
	for _, s := range n.current {
		if to, ok := n.trans[s][by]; ok {
			for _, t := range to {
				next[t] = true
			}
		}
	}
	n.current = make([]*State, 0, len(next)) // обновить текущее множество
	for s := range next {
		n.current = append(n.current, s)
	}
	return n.current
}

// Recognize запускает НКА по входной цепочке символов алфавита НКА и возвращает true, если цепочка принадлежит языку НКА, или false, если нет
func (n *NFA) Recognize(input string) bool {
	n.ResetCurrentStates()
	for _, c := range input {
		n.Transition(NewLetter(string(c)))
	}
	for _, s := range n.current {
		if s.IsTerminal() {
			return true
		}
	}
	return false
}

// RecognizeArray для массива Letter
func (n *NFA) RecognizeArray(input []*Letter) bool {

	n.ResetCurrentStates()
	for _, c := range input {
		n.Transition(c)
	}
	for _, s := range n.current {
		if s.IsTerminal() {
			return true
		}
	}
	return false
}
