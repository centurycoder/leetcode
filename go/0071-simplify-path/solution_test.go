package solution

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "example one",
			path: "/home/",
			want: "/home",
		},
		{
			name: "example two",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "example three",
			path: "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},
		{
			name: "parent at root",
			path: "/../",
			want: "/",
		},
		{
			name: "periods in directory names",
			path: "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
		{
			name: "root",
			path: "/",
			want: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Fatalf("simplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
