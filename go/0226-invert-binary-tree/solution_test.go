package solution

import "testing"

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "example one",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
			},
		},
		{
			name: "example two",
			root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
		},
		{
			name: "empty tree",
			want: nil,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 5},
			want: &TreeNode{Val: 5},
		},
		{
			name: "left-only chain",
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			want: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.root); !sameTree(got, tt.want) {
				t.Fatalf("invertTree() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func sameTree(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}

	return a.Val == b.Val && sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}
