package disposable

import "strings"

var black, white map[string]bool

// IsBlack checks if a given domain is one of disposable domains
func IsBlack(d string) bool {
	if black == nil {
		black = fill(Black)
	}
	return in(d, black)
}

// IsWhite checks if a given domain is one of domains that are treated
// as disposable but real actually.
func IsWhite(d string) bool {
	if white == nil {
		white = fill(White)
	}
	return in(d, white)
}

func fill(ls []string) map[string]bool {
	m := map[string]bool{}
	for _, i := range ls {
		m[i] = true
	}
	return m
}

func in(d string, m map[string]bool) bool {
	return m[strings.ToLower(d)]
}
