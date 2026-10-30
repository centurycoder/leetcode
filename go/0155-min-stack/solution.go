package solution

type minStackEntry struct {
	value int
	min   int
}

type MinStack struct {
	entries []minStackEntry
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(value int) {
	entry := minStackEntry{
		value: value,
		min:   value,
	}
	if len(s.entries) > 0 {
		minOnTop := s.entries[len(s.entries)-1].min
		if value > minOnTop {
			entry.min = minOnTop
		}
	}
	s.entries = append(s.entries, entry)
}

func (s *MinStack) Pop() {
	s.entries = s.entries[:len(s.entries)-1]
}

func (s *MinStack) Top() int {
	return s.entries[len(s.entries)-1].value
}

func (s *MinStack) GetMin() int {
	return s.entries[len(s.entries)-1].min
}
