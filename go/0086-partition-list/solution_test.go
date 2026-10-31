package solution

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		x      int
		want   []int
	}{
		{
			name:   "example one",
			values: []int{1, 4, 3, 2, 5, 2},
			x:      3,
			want:   []int{1, 2, 2, 4, 3, 5},
		},
		{
			name:   "example two",
			values: []int{2, 1},
			x:      2,
			want:   []int{1, 2},
		},
		{
			name: "empty list",
			x:    0,
		},
		{
			name:   "all values below pivot",
			values: []int{-3, -2, -1},
			x:      0,
			want:   []int{-3, -2, -1},
		},
		{
			name:   "all values at or above pivot",
			values: []int{0, 1, 2},
			x:      0,
			want:   []int{0, 1, 2},
		},
		{
			name:   "single node",
			values: []int{7},
			x:      7,
			want:   []int{7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listValues(partition(buildList(tt.values), tt.x))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("partition(%v, %d) = %v, want %v", tt.values, tt.x, got, tt.want)
			}
		})
	}
}

func TestPartitionPreservesNodeOrder(t *testing.T) {
	nodes := []*ListNode{
		{Val: 3},
		{Val: 1},
		{Val: 2},
		{Val: 1},
		{Val: 3},
		{Val: 2},
	}
	for i := 0; i+1 < len(nodes); i++ {
		nodes[i].Next = nodes[i+1]
	}

	got := listNodes(t, partition(nodes[0], 3))
	want := []*ListNode{nodes[1], nodes[3], nodes[2], nodes[5], nodes[0], nodes[4]}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("partition changed node identity or relative order: got %p, want %p", got, want)
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

func listNodes(t *testing.T, head *ListNode) []*ListNode {
	t.Helper()

	var nodes []*ListNode
	seen := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		if _, ok := seen[node]; ok {
			t.Fatal("partition result contains a cycle")
		}
		seen[node] = struct{}{}
		nodes = append(nodes, node)
	}
	return nodes
}
