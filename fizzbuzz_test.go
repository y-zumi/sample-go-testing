package tdd

import "testing"

func Test_Stringify(t *testing.T) {
	t.Run("when number is 1", func(t *testing.T) {
		if actual, expect := Stringify(1), "1"; actual != expect {
			t.Errorf("actual=%s, expect=%s", actual, expect)
		}
	})
	t.Run("when number is 2", func(t *testing.T) {
		if actual, expect := Stringify(2), "2"; actual != expect {
			t.Errorf("actual=%s, expect=%s", actual, expect)
		}
	})
	t.Run("when number is 100", func(t *testing.T) {
		if actual, expect := Stringify(100), "100"; actual != expect {
			t.Errorf("actual=%s, expect=%s", actual, expect)
		}
	})
}
