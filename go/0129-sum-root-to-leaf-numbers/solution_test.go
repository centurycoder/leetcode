package solution

import "testing"

func TestSumNumbers(t *testing.T) {
	deep := &TreeNode{Val: 1}
	current := deep
	for range 9 {
		current.Left = &TreeNode{Val: 0}
		current = current.Left
	}

	tests := []struct {
		name  string
		input string
		root  *TreeNode
		want  int
	}{
		{
			name:  "example one",
			input: "[1,2,3]",
			root:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want:  25,
		},
		{
			name:  "example two",
			input: "[4,9,0,5,1]",
			root: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
				Right: &TreeNode{Val: 0},
			},
			want: 1026,
		},
		{
			name:  "single nonzero digit",
			input: "[7]",
			root:  &TreeNode{Val: 7},
			want:  7,
		},
		{
			name:  "single zero digit",
			input: "[0]",
			root:  &TreeNode{Val: 0},
			want:  0,
		},
		{
			name:  "leading zeros",
			input: "[0,0,1,5]",
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 0, Left: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 1},
			},
			want: 6,
		},
		{
			name:  "right-only path with an interior zero",
			input: "[1,null,0,null,2]",
			root:  &TreeNode{Val: 1, Right: &TreeNode{Val: 0, Right: &TreeNode{Val: 2}}},
			want:  102,
		},
		{
			name:  "unequal path lengths share only their prefix",
			input: "[1,2,3,null,4]",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			},
			want: 137,
		},
		{
			name:  "ten-digit left-only path",
			input: "left chain: 1 -> 0 -> 0 -> 0 -> 0 -> 0 -> 0 -> 0 -> 0 -> 0",
			root:  deep,
			want:  1000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.root); got != tt.want {
				t.Fatalf("case %q (root=%s): sumNumbers() = %d, want %d", tt.name, tt.input, got, tt.want)
			}
		})
	}
}
