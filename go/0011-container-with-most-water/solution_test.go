package solution

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example one",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "example two",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "zero-height lines",
			height: []int{0, 0, 0},
			want:   0,
		},
		{
			name:   "equal heights use widest pair",
			height: []int{5, 5, 5, 5},
			want:   15,
		},
		{
			name:   "interior pair is optimal",
			height: []int{1, 2, 4, 3},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Fatalf("maxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
