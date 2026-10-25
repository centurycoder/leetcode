package solution

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "example one",
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			name:  "single zero",
			input: []int{0},
			want:  []int{0},
		},
		{
			name:  "all zeros",
			input: []int{0, 0, 0},
			want:  []int{0, 0, 0},
		},
		{
			name:  "no zeros",
			input: []int{-2, 5, 5},
			want:  []int{-2, 5, 5},
		},
		{
			name:  "leading and interspersed zeros",
			input: []int{0, 7, 0, -2, 0},
			want:  []int{7, -2, 0, 0, 0},
		},
		{
			name:  "zeros already at end",
			input: []int{1, 2, 0, 0},
			want:  []int{1, 2, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.input)
			moveZeroes(nums)
			if !slices.Equal(nums, tt.want) {
				t.Fatalf("moveZeroes() = %v, want %v", nums, tt.want)
			}
		})
	}
}
