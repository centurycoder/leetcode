package solution

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{
			name:  "example one",
			input: []int{-1, 0, 1, 2, -1, -4},
			want:  [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:  "example two",
			input: []int{0, 1, 1},
			want:  [][]int{},
		},
		{
			name:  "example three",
			input: []int{0, 0, 0},
			want:  [][]int{{0, 0, 0}},
		},
		{
			name:  "duplicate values produce one triplet",
			input: []int{-2, 0, 0, 2, 2},
			want:  [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canonicalTriplets(threeSum(slices.Clone(tt.input)))
			want := canonicalTriplets(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("threeSum() = %v, want %v", got, want)
			}
		})
	}
}

func canonicalTriplets(triplets [][]int) [][]int {
	canonical := make([][]int, len(triplets))
	for i, triplet := range triplets {
		canonical[i] = slices.Clone(triplet)
		slices.Sort(canonical[i])
	}
	sort.Slice(canonical, func(i, j int) bool {
		for k := range canonical[i] {
			if canonical[i][k] != canonical[j][k] {
				return canonical[i][k] < canonical[j][k]
			}
		}
		return len(canonical[i]) < len(canonical[j])
	})
	return canonical
}
