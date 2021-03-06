package comb

type State struct {
	source         []rune
	line, position int
}

func NewState(source string) *State {
	return &State{source: ([]rune)(source)}
}

func (s State) exhausted() bool {
	return s.position >= len(s.source)
}

func (s State) currentRune() rune {
	if s.exhausted() {
		return '\x00'
	}

	return s.source[s.position]
}

func (s *State) increment() {
	if s.currentRune() == '\n' {
		s.line++
	}

	s.position++
}
