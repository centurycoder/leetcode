package solution

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "example one",
			input: []int{100, 4, 200, 1, 3, 2},
			want:  4,
		},
		{
			name:  "example two",
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want:  9,
		},
		{
			name:  "example three with duplicates",
			input: []int{1, 0, 1, 2},
			want:  3,
		},
		{
			name:  "empty input",
			input: []int{},
			want:  0,
		},
		{
			name:  "negative consecutive values",
			input: []int{-4, -1, -3, -2, 5},
			want:  4,
		},
		{
			name:  "duplicate values do not extend sequence",
			input: []int{7, 7, 7, 8, 8, 9},
			want:  3,
		},
		{
			name:  "sequences separated by gaps",
			input: []int{10, 11, 12, 100, 101},
			want:  3,
		},
		{
			name: "values near allowed range limits",
			input: []int{
				1_000_000_000, -999_999_997, 0,
				-1_000_000_000, 999_999_999, -999_999_999,
				999_999_997, -999_999_998, 999_999_998,
				999_999_996,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.input); got != tt.want {
				t.Errorf("longestConsecutive() = %d, want %d", got, tt.want)
			}
		})
	}
}
