package solution

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "duplicate runs in the middle",
			values: []int{1, 2, 3, 3, 4, 4, 5},
			want:   []int{1, 2, 5},
		},
		{
			name:   "duplicate run at the head",
			values: []int{1, 1, 1, 2, 3},
			want:   []int{2, 3},
		},
		{
			name:   "duplicate run at the tail",
			values: []int{1, 2, 2},
			want:   []int{1},
		},
		{
			name:   "all values duplicated",
			values: []int{1, 1, 2, 2},
			want:   nil,
		},
		{
			name:   "no duplicates",
			values: []int{-2, 0, 3},
			want:   []int{-2, 0, 3},
		},
		{
			name:   "single node",
			values: []int{7},
			want:   []int{7},
		},
		{
			name: "empty list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listValues(deleteDuplicates(buildList(tt.values)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("deleteDuplicates(%v) = %v, want %v", tt.values, got, tt.want)
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
