package solution

import "testing"

func TestFlatten(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n1.Left = n2
	n1.Right = n5
	n2.Left = n3
	n2.Right = n4
	n5.Right = n6

	single := &TreeNode{Val: 0}
	tests := []struct {
		name string
		root *TreeNode
		want []*TreeNode
	}{
		{
			name: "example one",
			root: n1,
			want: []*TreeNode{n1, n2, n3, n4, n5, n6},
		},
		{
			name: "empty tree",
		},
		{
			name: "single node",
			root: single,
			want: []*TreeNode{single},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)

			current := tt.root
			for i, want := range tt.want {
				if current == nil {
					t.Fatalf("flatten() chain ended at index %d; want node with value %d", i, want.Val)
				}
				if current != want {
					t.Fatalf("flatten() node at index %d = %p (value %d), want original node %p (value %d)", i, current, current.Val, want, want.Val)
				}
				if current.Left != nil {
					t.Fatalf("flatten() node at index %d (value %d) has non-nil Left", i, current.Val)
				}
				current = current.Right
			}

			if current != nil {
				t.Fatalf("flatten() chain has an unexpected node after %d expected nodes: %p (value %d)", len(tt.want), current, current.Val)
			}
		})
	}
}
