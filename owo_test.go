package owo

import "testing"

func TestTranslate(t *testing.T) {
	want := "I haz nu mouth and I must scweam"
	if got := Translate("I have no mouth and I must scream", false, false); got != want {
		t.Errorf("Translate() = %q, want %q", got, want)
	}
}