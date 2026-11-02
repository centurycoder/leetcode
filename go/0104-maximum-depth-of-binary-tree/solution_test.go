package solution

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example one",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "example two",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: 2,
		},
		{
			name: "empty tree",
			want: 0,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 7},
			want: 1,
		},
		{
			name: "depth follows longest branch",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{Val: 5},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Fatalf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
