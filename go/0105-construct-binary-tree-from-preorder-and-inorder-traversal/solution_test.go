package solution

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			name:     "example one",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
		},
		{
			name:     "single negative value",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     &TreeNode{Val: -1},
		},
		{
			name:     "left-skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			want:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
		},
		{
			name:     "right-skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		},
		{
			name: "empty traversals",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("buildTree() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
