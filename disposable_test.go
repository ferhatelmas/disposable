package disposable

import "testing"

func TestIsBlack(t *testing.T) {
	domains := []string{"0-mail.com", "1CE.US"}
	test(domains, true, IsBlack, t)
}

func TestIsWhite(t *testing.T) {
	domains := []string{"fasTmaIl.com", "antiChef.com"}
	test(domains, true, IsWhite, t)
}

func TestNoWhiteIsBlack(t *testing.T) {
	test(White, false, IsBlack, t)
}

func TestAllBlackIsBlack(t *testing.T) {
	test(Black, true, IsBlack, t)
}

func TestNoBlackIsWhite(t *testing.T) {
	test(Black, false, IsWhite, t)
}

func TestAllWhiteIsWhite(t *testing.T) {
	test(White, true, IsWhite, t)
}

func TestNeitherBlackNorWhite(t *testing.T) {
	domains := []string{"google.com", "yahoo.com"}
	test(domains, false, IsBlack, t)
	test(domains, false, IsWhite, t)
}

func test(domains []string, expected bool, check func(string) bool, t *testing.T) {
	for i, d := range domains {
		if res := check(d); res != expected {
			t.Errorf("%d: %q - expected %t, got %t", i, d, expected, res)
		}
	}
}

func TestNoDuplicate(t *testing.T) {
	checkDuplicates(Black, t)
	checkDuplicates(White, t)
}

func checkDuplicates(ls []string, t *testing.T) {
	m := map[string]bool{}
	for _, s := range ls {
		if m[s] {
			t.Fatalf(s + " is a duplicate.")
		} else {
			m[s] = true
		}
	}
}
