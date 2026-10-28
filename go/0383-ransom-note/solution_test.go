package solution

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			name:       "example one",
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			name:       "example two",
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			name:       "example three",
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Fatalf("canConstruct(%q, %q) = %t, want %t", tt.ransomNote, tt.magazine, got, tt.want)
			}
		})
	}
}
