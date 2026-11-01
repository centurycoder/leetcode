package solution

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		left   int
		right  int
		want   []int
	}{
		{
			name:   "reverse middle section",
			values: []int{1, 2, 3, 4, 5},
			left:   2,
			right:  4,
			want:   []int{1, 4, 3, 2, 5},
		},
		{
			name:   "single node",
			values: []int{5},
			left:   1,
			right:  1,
			want:   []int{5},
		},
		{
			name:   "reverse from head",
			values: []int{1, 2, 3, 4},
			left:   1,
			right:  3,
			want:   []int{3, 2, 1, 4},
		},
		{
			name:   "reverse through tail",
			values: []int{1, 2, 3, 4},
			left:   2,
			right:  4,
			want:   []int{1, 4, 3, 2},
		},
		{
			name:   "reverse entire list",
			values: []int{1, 2, 3},
			left:   1,
			right:  3,
			want:   []int{3, 2, 1},
		},
		{
			name:   "unchanged range",
			values: []int{1, 2, 3},
			left:   2,
			right:  2,
			want:   []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listValues(reverseBetween(buildList(tt.values), tt.left, tt.right))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("reverseBetween(%v, %d, %d) = %v, want %v", tt.values, tt.left, tt.right, got, tt.want)
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
