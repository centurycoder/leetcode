package solution

import "testing"

func TestIsSubsequence(t *testing.T) {

	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "test-1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "test-2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.s, tt.t); got != tt.want {
				t.Fatalf("test %s expect %t, got: %t", tt.name, tt.want, got)
			}
		})
	}
}
