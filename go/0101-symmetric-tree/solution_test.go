package solution

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name  string
		input string
		root  *TreeNode
		want  bool
	}{
		{
			name: "example one is symmetric",
			input: "[1,2,2,3,4,4,3]",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
			},
			want: true,
		},
		{
			name: "example two is not symmetric",
			input: "[1,2,2,null,3,null,3]",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
			want: false,
		},
		{
			name: "single node is symmetric",
			input: "[-100]",
			root: &TreeNode{Val: -100},
			want: true,
		},
		{
			name: "mirrored shape with unequal values is not symmetric",
			input: "[0,-1,-1,null,2,3]",
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -1, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: -1, Left: &TreeNode{Val: 3}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.root); got != tt.want {
				t.Fatalf("case %q (root=%s): isSymmetric() = %t, want %t", tt.name, tt.input, got, tt.want)
			}
		})
	}
}
