package tdd

import "testing"

func Test_Stringify(t *testing.T) {
	if actual, expect := Stringify(1), "1"; actual != expect {
		t.Errorf("actual=%s, expect=%s", actual, expect)
	}
}
