package durafmt

import "testing"

func TestInternationalStringZero(t *testing.T) {
	if got := Parse(0).InternationalString(); got != "0 s" {
		t.Errorf("Parse(0).InternationalString() = %q, want 0 s", got)
	}
	for _, tc := range []struct{ input, want string }{
		{"0h", "0 h"}, {"0m", "0 m"}, {"0s", "0 s"},
		{"0ms", "0 ms"}, {"0µs", "0 µs"}, {"-0s", "-0 s"},
	} {
		t.Run(tc.input, func(t *testing.T) {
			d, err := ParseString(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := d.InternationalString(); got != tc.want {
				t.Errorf("InternationalString() = %q, want %q", got, tc.want)
			}
		})
	}
}
