package solution

import "testing"

func TestConnect(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n7 := &Node{Val: 7}
	single := &Node{Val: 42}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Right = n7

	tests := []struct {
		name     string
		root     *Node
		wantNext map[*Node]*Node
	}{
		{
			name: "sparse tree",
			root: n1,
			wantNext: map[*Node]*Node{
				n1: nil,
				n2: n3,
				n3: nil,
				n4: n5,
				n5: n7,
				n7: nil,
			},
		},
		{
			name:     "empty tree",
			root:     nil,
			wantNext: map[*Node]*Node{},
		},
		{
			name:     "single node",
			root:     single,
			wantNext: map[*Node]*Node{single: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.root); got != tt.root {
				t.Fatalf("connect() = %p, want original root %p", got, tt.root)
			}
			for node, want := range tt.wantNext {
				if node.Next != want {
					t.Errorf("node %d Next = %p, want %p", node.Val, node.Next, want)
				}
			}
		})
	}
}
