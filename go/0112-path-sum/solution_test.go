package solution

import "testing"

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		root      *TreeNode
		targetSum int
		want      bool
	}{
		{
			name:  "example one has a matching root-to-leaf path",
			input: "[5,4,8,11,null,13,4,7,2,null,null,null,1]",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 13},
					Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}},
				},
			},
			targetSum: 22,
			want:      true,
		},
		{
			name:      "example two has no matching path",
			input:     "[1,2,3]",
			root:      &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			targetSum: 5,
			want:      false,
		},
		{
			name:      "example three is an empty tree",
			input:     "[]",
			targetSum: 0,
			want:      false,
		},
		{
			name:      "matching sum at a non-leaf is not a path",
			input:     "[1,2,3]",
			root:      &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			targetSum: 3,
			want:      false,
		},
		{
			name:      "negative values can form the target sum",
			input:     "[-2,null,-3]",
			root:      &TreeNode{Val: -2, Right: &TreeNode{Val: -3}},
			targetSum: -5,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.root, tt.targetSum); got != tt.want {
				t.Fatalf("case %q (root=%s, targetSum=%d): hasPathSum() = %t, want %t", tt.name, tt.input, tt.targetSum, got, tt.want)
			}
		})
	}
}
