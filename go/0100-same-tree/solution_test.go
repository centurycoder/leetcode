package solution

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "example one",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: true,
		},
		{
			name: "example two has different structure",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:    &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			name: "example three has different values",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			name: "both trees are empty",
			want: true,
		},
		{
			name: "one tree is empty",
			p:    &TreeNode{Val: 1},
			want: false,
		},
		{
			name: "same values and structure with distinct nodes",
			p:    &TreeNode{Val: 4, Left: &TreeNode{Val: -2}},
			q:    &TreeNode{Val: 4, Left: &TreeNode{Val: -2}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.p, tt.q); got != tt.want {
				t.Fatalf("isSameTree() = %t, want %t", got, tt.want)
			}
		})
	}
}
