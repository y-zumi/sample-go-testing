package tdd

import "testing"

type TestCase struct {
	description string
	input       int
	output      string
}

func Test_Stringify(t *testing.T) {
	t.Run("normal number", func(t *testing.T) {
		cases := []TestCase{
			{"number is 1", 1, "1"},
			{"number is 101", 101, "101"},
		}
		for _, c := range cases {
			t.Run(c.description, func(t *testing.T) {
				if actual, expect := Stringify(c.input), c.output; actual != expect {
					t.Errorf("actual=%s, expect=%s", actual, expect)
				}
			})
		}
	})

	t.Run("Multiple of 3", func(t *testing.T) {
		cases := []TestCase{
			{"number is 3", 3, "fizz"},
			{"number is 99", 99, "fizz"},
		}
		for _, c := range cases {
			t.Run(c.description, func(t *testing.T) {
				if actual, expect := Stringify(c.input), c.output; actual != expect {
					t.Errorf("actual=%s, expect=%s", actual, expect)
				}
			})
		}
	})

	t.Run("Multiple of 5", func(t *testing.T) {
		cases := []TestCase{
			{"number is 5", 5, "buzz"},
			{"number is 100", 100, "buzz"},
		}
		for _, c := range cases {
			t.Run(c.description, func(t *testing.T) {
				if actual, expect := Stringify(c.input), c.output; actual != expect {
					t.Errorf("actual=%s, expect=%s", actual, expect)
				}
			})
		}
	})

	t.Run("Multiple of 15", func(t *testing.T) {
		cases := []TestCase{
			{"number is 15", 15, "fizzbuzz"},
			{"number is 105", 105, "fizzbuzz"},
		}
		for _, c := range cases {
			t.Run(c.description, func(t *testing.T) {
				if actual, expect := Stringify(c.input), c.output; actual != expect {
					t.Errorf("actual=%s, expect=%s", actual, expect)
				}
			})
		}
	})

}
