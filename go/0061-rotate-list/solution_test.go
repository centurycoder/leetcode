package solution

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		k      int
		want   []int
	}{
		{
			name:   "example one",
			values: []int{1, 2, 3, 4, 5},
			k:      2,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			name:   "example two",
			values: []int{0, 1, 2},
			k:      4,
			want:   []int{2, 0, 1},
		},
		{
			name: "empty list",
			k:    3,
		},
		{
			name:   "single node",
			values: []int{7},
			k:      5,
			want:   []int{7},
		},
		{
			name:   "zero rotations",
			values: []int{1, 2, 3},
			k:      0,
			want:   []int{1, 2, 3},
		},
		{
			name:   "rotations equal list length",
			values: []int{1, 2, 3},
			k:      3,
			want:   []int{1, 2, 3},
		},
		{
			name:   "large rotation count",
			values: []int{1, 2, 3},
			k:      2000000000,
			want:   []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listValues(rotateRight(buildList(tt.values), tt.k))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("rotateRight(%v, %d) = %v, want %v", tt.values, tt.k, got, tt.want)
			}
		})
	}
}

func buildList(values []int) *ListNode {
	var head, tail *ListNode
	for _, value := range values {
		node := &ListNode{Val: value}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func listValues(head *ListNode) []int {
	var values []int
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}
