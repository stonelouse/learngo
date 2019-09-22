package stringutil

import "strings"

// SplitIn2 returns ... and the linter is satisfied
// Named return values
func MySplitIn2(value, delimiter string) (first string, last string) {
	s := strings.Split(value, delimiter)
	first = s[0]
	last = s[1]
	return // ... "naked" return
}
