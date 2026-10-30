package solution

import "testing"

func TestMinStackOperations(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if got := stack.GetMin(); got != -3 {
		t.Fatalf("GetMin() = %d, want -3", got)
	}

	stack.Pop()
	if got := stack.Top(); got != 0 {
		t.Fatalf("Top() = %d, want 0", got)
	}
	if got := stack.GetMin(); got != -2 {
		t.Fatalf("GetMin() after Pop() = %d, want -2", got)
	}
}

func TestMinStackDuplicateMinimum(t *testing.T) {
	stack := Constructor()
	stack.Push(2)
	stack.Push(1)
	stack.Push(1)

	stack.Pop()
	if got := stack.GetMin(); got != 1 {
		t.Fatalf("GetMin() after popping duplicate minimum = %d, want 1", got)
	}

	stack.Pop()
	if got := stack.GetMin(); got != 2 {
		t.Fatalf("GetMin() after removing minimum = %d, want 2", got)
	}
}

func TestMinStackInt32Bounds(t *testing.T) {
	const (
		minInt32 = -1 << 31
		maxInt32 = 1<<31 - 1
	)

	stack := Constructor()
	stack.Push(maxInt32)
	stack.Push(minInt32)
	if got := stack.GetMin(); got != minInt32 {
		t.Fatalf("GetMin() = %d, want %d", got, minInt32)
	}

	stack.Pop()
	if got := stack.GetMin(); got != maxInt32 {
		t.Fatalf("GetMin() after Pop() = %d, want %d", got, maxInt32)
	}
}
