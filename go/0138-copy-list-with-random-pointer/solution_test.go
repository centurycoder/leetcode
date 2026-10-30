package solution

import "testing"

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name          string
		values        []int
		randomIndices []int
	}{
		{
			name:          "example one",
			values:        []int{7, 13, 11, 10, 1},
			randomIndices: []int{-1, 0, 4, 2, 0},
		},
		{
			name:          "example two",
			values:        []int{1, 2},
			randomIndices: []int{1, 1},
		},
		{
			name:          "duplicate values and nil random pointers",
			values:        []int{3, 3, 3},
			randomIndices: []int{-1, 0, -1},
		},
		{
			name:          "empty list",
			values:        nil,
			randomIndices: nil,
		},
		{
			name:          "single node points to itself",
			values:        []int{5},
			randomIndices: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originals := buildList(tt.values, tt.randomIndices)
			copiedHead := copyRandomList(headOf(originals))
			assertDeepCopy(t, originals, copiedHead, tt.randomIndices)
		})
	}
}

func buildList(values, randomIndices []int) []*Node {
	nodes := make([]*Node, len(values))
	for i, value := range values {
		nodes[i] = &Node{Val: value}
	}
	for i, node := range nodes {
		if i+1 < len(nodes) {
			node.Next = nodes[i+1]
		}
		if randomIndices[i] >= 0 {
			node.Random = nodes[randomIndices[i]]
		}
	}
	return nodes
}

func headOf(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

func assertDeepCopy(t *testing.T, originals []*Node, copiedHead *Node, randomIndices []int) {
	t.Helper()

	if len(originals) == 0 {
		if copiedHead != nil {
			t.Fatalf("copyRandomList(nil) = %p, want nil", copiedHead)
		}
		return
	}

	copies := make([]*Node, len(originals))
	current := copiedHead
	for i := range copies {
		if current == nil {
			t.Fatalf("copied list ended at index %d; want %d nodes", i, len(copies))
		}
		if current == originals[i] {
			t.Fatalf("copied node %d aliases the original node", i)
		}
		copies[i] = current
		current = current.Next
	}
	if current != nil {
		t.Fatalf("copied list has more than %d nodes", len(copies))
	}

	for i, original := range originals {
		if copies[i].Val != original.Val {
			t.Errorf("copy %d value = %d, want %d", i, copies[i].Val, original.Val)
		}
		var wantNext *Node
		if i+1 < len(copies) {
			wantNext = copies[i+1]
		}
		if copies[i].Next != wantNext {
			t.Errorf("copy %d next = %p, want %p", i, copies[i].Next, wantNext)
		}
		var wantRandom *Node
		if randomIndices[i] >= 0 {
			wantRandom = copies[randomIndices[i]]
		}
		if copies[i].Random != wantRandom {
			t.Errorf("copy %d random = %p, want %p", i, copies[i].Random, wantRandom)
		}
	}
}
