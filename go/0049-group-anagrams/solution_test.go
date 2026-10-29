package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "example one",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:  "example two empty string",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			name:  "example three single string",
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
		{
			name:  "repeated strings and repeated empty strings",
			input: []string{"", "a", "", "a"},
			want:  [][]string{{"", ""}, {"a", "a"}},
		},
		{
			name:  "different letter frequencies form different groups",
			input: []string{"ab", "ba", "abc", "cab", "ac"},
			want:  [][]string{{"ab", "ba"}, {"abc", "cab"}, {"ac"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canonicalGroups(groupAnagrams(tt.input))
			want := canonicalGroups(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("groupAnagrams() = %v, want %v", got, want)
			}
		})
	}
}

func canonicalGroups(groups [][]string) [][]string {
	canonical := make([][]string, len(groups))
	for i, group := range groups {
		canonical[i] = append([]string(nil), group...)
		sort.Strings(canonical[i])
	}
	sort.Slice(canonical, func(i, j int) bool {
		for k := 0; k < len(canonical[i]) && k < len(canonical[j]); k++ {
			if canonical[i][k] != canonical[j][k] {
				return canonical[i][k] < canonical[j][k]
			}
		}
		return len(canonical[i]) < len(canonical[j])
	})
	return canonical
}
