package text

import "testing"

func TestTextChecker_CheckOk(t *testing.T) {
	checker := TextChecker{"ok"}

	assertCheck(t, checker, true)
}

func TestTextChecker_CheckNotOk(t *testing.T) {
	checker := TextChecker{""}

	assertCheck(t, checker, false)
}

func assertCheck(t *testing.T, checker TextChecker, want bool) {
	if checker.Check() != want {
		t.Errorf("got %t want %t", checker.Check(), want)
	}
}
