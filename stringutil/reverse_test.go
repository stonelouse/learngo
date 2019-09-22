package stringutil

// run D:\NoScan\home.rus\dev.ext\ws.go>go test github.com/stonelouse/learngo/stringutil

// lightweight test framework composed of the 'go test' command
import "testing"

// ... test framework run each function with name pattern TestXXX
func TestReverse(t *testing.T) {
	cases :=
		[]struct {
			in, want string
		}{
			{"Hello world", "dlrow olleH"},
			{"", ""},
			{"Hello, 世界", "界世 ,olleH"},
		}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
