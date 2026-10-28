package solution

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := []struct {
		name    string
		numbers []int
		target  int
		expect  []int
	}{
		{
			name:    "t1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			expect:  []int{1, 2},
		},
		{
			name:    "t2",
			numbers: []int{5, 25, 75},
			target:  100,
			expect:  []int{2, 3},
		},
	}

	for _, tt := range result {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Fatalf("tesing %s failed: expect %v, got %v", tt.name, tt.expect, got)
			}
		})
	}
}
