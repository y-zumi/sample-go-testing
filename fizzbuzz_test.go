package tdd

import "testing"

func Test_Stringify(t *testing.T) {
	t.Run("normal number", func(t *testing.T) {
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
		t.Run("when number is 101", func(t *testing.T) {
			if actual, expect := Stringify(101), "101"; actual != expect {
				t.Errorf("actual=%s, expect=%s", actual, expect)
			}
		})
	})

	t.Run("Multiple of 3", func(t *testing.T) {
		t.Run("when number is 3", func(t *testing.T) {
			if actual, expect := Stringify(3), "fizz"; actual != expect {
				t.Errorf("actual=%s, expect=%s", actual, expect)
			}
		})
		t.Run("when number is 99", func(t *testing.T) {
			if actual, expect := Stringify(99), "fizz"; actual != expect {
				t.Errorf("actual=%s, expect=%s", actual, expect)
			}
		})
	})

	t.Run("Multiple of 5", func(t *testing.T) {
		t.Run("when number is 5", func(t *testing.T) {
			if actual, expect := Stringify(5), "buzz"; actual != expect {
				t.Errorf("actual=%s, expect=%s", actual, expect)
			}
		})
		t.Run("when number is 100", func(t *testing.T) {
			if actual, expect := Stringify(100), "buzz"; actual != expect {
				t.Errorf("actual=%s, expect=%s", actual, expect)
			}
		})
	})

}
