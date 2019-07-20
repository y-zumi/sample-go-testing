package tdd

import "testing"

func Test_Stringify(t *testing.T) {
	Stringify()
	if actual, expect := Stringify(), "0"; actual != expect {
		t.Errorf("actual=%s, expect=%s", actual, expect)
	}
}
