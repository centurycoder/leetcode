package solution

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "example one",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "example two",
			input: "race a car",
			want:  false,
		},
		{
			name:  "example three",
			input: " ",
			want:  true,
		},
		{
			name:  "punctuation-only string",
			input: ".,!?:",
			want:  true,
		},
		{
			name:  "digits are alphanumeric",
			input: "1a2A1",
			want:  true,
		},
		{
			name:  "case is ignored",
			input: "AbBa",
			want:  true,
		},
		{
			name:  "mismatch after filtering",
			input: "0P",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.input); got != tt.want {
				t.Fatalf("isPalindrome(%q) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}
}
