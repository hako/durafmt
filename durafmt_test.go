package durafmt

import (
	"fmt"
	"testing"
	"time"
)

var (
	testStrings []struct {
		test     string
		expected string
	}
	testTimes []struct {
		test     time.Duration
		expected string
	}
	testTimesWithLimitUnit []struct {
		test      time.Duration
		limitUnit string
		expected  string
	}
	testTimesWithLimit []struct {
		test     time.Duration
		limitN   int
		expected string
	}
)

// TestParse for durafmt time.Duration conversion.
func TestParse(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Microsecond, "1 microsecond"},
		{1 * time.Millisecond, "1 millisecond"},
		{1 * time.Second, "1 second"},
		{1 * time.Hour, "1 hour"},
		{1 * time.Minute, "1 minute"},
		{2 * time.Microsecond, "2 microseconds"},
		{2 * time.Millisecond, "2 milliseconds"},
		{2 * time.Second, "2 seconds"},
		{2 * time.Minute, "2 minutes"},
		{1 * time.Hour, "1 hour"},
		{2 * time.Hour, "2 hours"},
		{10 * time.Hour, "10 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{120 * time.Hour, "5 days"},
		{168 * time.Hour, "1 week"},
		{672 * time.Hour, "4 weeks"},
		{8759 * time.Hour, "52 weeks 23 hours"},
		{8760 * time.Hour, "1 year"},
		{17519 * time.Hour, "1 year 52 weeks 23 hours"},
		{17520 * time.Hour, "2 years"},
		{26279 * time.Hour, "2 years 52 weeks 23 hours"},
		{26280 * time.Hour, "3 years"},
		{201479 * time.Hour, "22 years 52 weeks 23 hours"},
		{201480 * time.Hour, "23 years"},
		{-1 * time.Second, "-1 second"},
		{-10 * time.Second, "-10 seconds"},
		{-100 * time.Second, "-1 minute 40 seconds"},
		{-1 * time.Millisecond, "-1 millisecond"},
		{-10 * time.Millisecond, "-10 milliseconds"},
		{-100 * time.Millisecond, "-100 milliseconds"},
		{-1 * time.Microsecond, "-1 microsecond"},
		{-10 * time.Microsecond, "-10 microseconds"},
		{-100 * time.Microsecond, "-100 microseconds"},
		{-1000 * time.Microsecond, "-1 millisecond"},
		{-1000000 * time.Microsecond, "-1 second"},
		{-1001000 * time.Microsecond, "-1 second 1 millisecond"},
		{-1010000 * time.Microsecond, "-1 second 10 milliseconds"},
		{-1001001 * time.Microsecond, "-1 second 1 millisecond 1 microsecond"},
		{-1001002 * time.Microsecond, "-1 second 1 millisecond 2 microseconds"},
	}

	for _, table := range testTimes {
		result := Parse(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q",
				table.test, result, result, table.expected)
		}
	}
}

func TestParseWithLimitToUnit(t *testing.T) {
	testTimesWithLimitUnit = []struct {
		test      time.Duration
		limitUnit string
		expected  string
	}{
		{87593183 * time.Second, "seconds", "87593183 seconds"},
		{87593183 * time.Second, "minutes", "1459886 minutes 23 seconds"},
		{87593183 * time.Second, "hours", "24331 hours 26 minutes 23 seconds"},
		{87593183 * time.Second, "days", "1013 days 19 hours 26 minutes 23 seconds"},
		{87593183 * time.Second, "weeks", "144 weeks 5 days 19 hours 26 minutes 23 seconds"},
		{87593183 * time.Second, "years", "2 years 40 weeks 3 days 19 hours 26 minutes 23 seconds"},
		{87593183 * time.Second, "", "2 years 40 weeks 3 days 19 hours 26 minutes 23 seconds"},
	}

	for _, table := range testTimesWithLimitUnit {
		result := Parse(table.test).LimitToUnit(table.limitUnit).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q",
				table.test, result, result, table.expected)
		}
	}
}

func TestParseWithLimitN(t *testing.T) {
	testTimesWithLimit = []struct {
		test     time.Duration
		limitN   int
		expected string
	}{
		{1 * time.Millisecond, 0, "1 millisecond"},
		{8759 * time.Hour, 0, "52 weeks 23 hours"},
		{17519 * time.Hour, 0, "1 year 52 weeks 23 hours"},
		{-1 * time.Second, 0, "-1 second"},
		{-100 * time.Second, 0, "-1 minute 40 seconds"},
		{1 * time.Millisecond, 1, "1 millisecond"},
		{8759 * time.Hour, 1, "52 weeks"},
		{17519 * time.Hour, 1, "1 year"},
		{-1 * time.Second, 1, "-1 second"},
		{-100 * time.Second, 1, "-1 minute"},
		{1 * time.Millisecond, 2, "1 millisecond"},
		{8759 * time.Hour, 2, "52 weeks 23 hours"},
		{17519 * time.Hour, 2, "1 year 52 weeks"},
		{-1 * time.Second, 2, "-1 second"},
		{-100 * time.Second, 2, "-1 minute 40 seconds"},
		{1 * time.Millisecond, 3, "1 millisecond"},
		{8759 * time.Hour, 3, "52 weeks 23 hours"},
		{17519 * time.Hour, 3, "1 year 52 weeks 23 hours"},
		{-1 * time.Second, 3, "-1 second"},
		{-100 * time.Second, 3, "-1 minute 40 seconds"},
	}

	for _, table := range testTimesWithLimit {
		result := Parse(table.test).LimitFirstN(table.limitN).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q",
				table.test, result, result, table.expected)
		}
	}
}

// TestParseShort for durafmt time.Duration conversion, short version.
func TestParseShort(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Microsecond, "1 microsecond"},
		{1 * time.Millisecond, "1 millisecond"},
		{1 * time.Second, "1 second"},
		{1 * time.Hour, "1 hour"},
		{1 * time.Minute, "1 minute"},
		{2 * time.Microsecond, "2 microseconds"},
		{2 * time.Millisecond, "2 milliseconds"},
		{2 * time.Second, "2 seconds"},
		{2 * time.Minute, "2 minutes"},
		{1 * time.Hour, "1 hour"},
		{2 * time.Hour, "2 hours"},
		{10 * time.Hour, "10 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{120 * time.Hour, "5 days"},
		{168 * time.Hour, "1 week"},
		{672 * time.Hour, "4 weeks"},
		{8759 * time.Hour, "52 weeks"},
		{8760 * time.Hour, "1 year"},
		{17519 * time.Hour, "1 year"},
		{17520 * time.Hour, "2 years"},
		{26279 * time.Hour, "2 years"},
		{26280 * time.Hour, "3 years"},
		{201479 * time.Hour, "22 years"},
		{201480 * time.Hour, "23 years"},
		{-1 * time.Second, "-1 second"},
		{-10 * time.Second, "-10 seconds"},
		{-100 * time.Second, "-1 minute"},
		{-1 * time.Millisecond, "-1 millisecond"},
		{-10 * time.Millisecond, "-10 milliseconds"},
		{-100 * time.Millisecond, "-100 milliseconds"},
		{-1 * time.Microsecond, "-1 microsecond"},
		{-10 * time.Microsecond, "-10 microseconds"},
		{-100 * time.Microsecond, "-100 microseconds"},
		{-1000 * time.Microsecond, "-1 millisecond"},
		{-1000000 * time.Microsecond, "-1 second"},
		{-1001000 * time.Microsecond, "-1 second"},
		{-1010000 * time.Microsecond, "-1 second"},
		{-1001001 * time.Microsecond, "-1 second"},
		{-1001002 * time.Microsecond, "-1 second"},
	}

	for _, table := range testTimes {
		result := ParseShort(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q",
				table.test, result, result, table.expected)
		}
	}
}

// TestParseString for durafmt duration string conversion.
func TestParseString(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1µs", "1 microsecond"},
		{"2µs", "2 microseconds"},
		{"1ms", "1 millisecond"},
		{"2ms", "2 milliseconds"},
		{"1s", "1 second"},
		{"2s", "2 seconds"},
		{"1m", "1 minute"},
		{"2m", "2 minutes"},
		{"1h", "1 hour"},
		{"2h", "2 hours"},
		{"10h", "10 hours"},
		{"24h", "1 day"},
		{"48h", "2 days"},
		{"120h", "5 days"},
		{"168h", "1 week"},
		{"672h", "4 weeks"},
		{"8759h", "52 weeks 23 hours"},
		{"8760h", "1 year"},
		{"17519h", "1 year 52 weeks 23 hours"},
		{"17520h", "2 years"},
		{"26279h", "2 years 52 weeks 23 hours"},
		{"26280h", "3 years"},
		{"201479h", "22 years 52 weeks 23 hours"},
		{"201480h", "23 years"},
		{"1m0s", "1 minute"},
		{"1m2s", "1 minute 2 seconds"},
		{"3h4m5s", "3 hours 4 minutes 5 seconds"},
		{"6h7m8s9ms", "6 hours 7 minutes 8 seconds 9 milliseconds"},
		{"0µs", "0 microseconds"},
		{"0ms", "0 milliseconds"},
		{"0s", "0 seconds"},
		{"0m", "0 minutes"},
		{"0h", "0 hours"},
		{"0m1µs", "1 microsecond"},
		{"0m1ms1µs", "1 millisecond 1 microsecond"},
		{"0m1ms", "1 millisecond"},
		{"0m1s", "1 second"},
		{"0m1m", "1 minute"},
		{"0m2ms", "2 milliseconds"},
		{"0m2s", "2 seconds"},
		{"0m2m", "2 minutes"},
		{"0m2m3h", "3 hours 2 minutes"},
		{"0m2m34h", "1 day 10 hours 2 minutes"},
		{"0m56h7m8ms", "2 days 8 hours 7 minutes 8 milliseconds"},
		{"-1µs", "-1 microsecond"},
		{"-1ms", "-1 millisecond"},
		{"-1s", "-1 second"},
		{"-1m", "-1 minute"},
		{"-1h", "-1 hour"},
		{"-2µs", "-2 microseconds"},
		{"-2ms", "-2 milliseconds"},
		{"-2s", "-2 seconds"},
		{"-2m", "-2 minutes"},
		{"-2h", "-2 hours"},
		{"-10h", "-10 hours"},
		{"-24h", "-1 day"},
		{"-48h", "-2 days"},
		{"-120h", "-5 days"},
		{"-168h", "-1 week"},
		{"-672h", "-4 weeks"},
		{"-8760h", "-1 year"},
		{"-1m0s", "-1 minute"},
		{"-0m2s", "-2 seconds"},
		{"-0m2m", "-2 minutes"},
		{"-0m2m3h", "-3 hours 2 minutes"},
		{"-0m2m34h", "-1 day 10 hours 2 minutes"},
		{"-0µs", "-0 microseconds"},
		{"-0ms", "-0 milliseconds"},
		{"-0s", "-0 seconds"},
		{"-0m", "-0 minutes"},
		{"-0h", "-0 hours"},
	}

	for _, table := range testStrings {
		d, err := ParseString(table.test)
		if err != nil {
			t.Errorf("%q", err)
		}
		result := d.String()
		if result != table.expected {
			t.Errorf("d.String() = %q. got %q, expected %q",
				table.test, result, table.expected)
		}
	}
}

// TestParseStringShort for durafmt duration string conversion, short version.
func TestParseStringShort(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1µs", "1 microsecond"},
		{"1ms", "1 millisecond"},
		{"2µs", "2 microseconds"},
		{"2ms", "2 milliseconds"},
		{"1s", "1 second"},
		{"2s", "2 seconds"},
		{"1m", "1 minute"},
		{"2m", "2 minutes"},
		{"1h", "1 hour"},
		{"2h", "2 hours"},
		{"10h", "10 hours"},
		{"24h", "1 day"},
		{"48h", "2 days"},
		{"120h", "5 days"},
		{"168h", "1 week"},
		{"672h", "4 weeks"},
		{"8759h", "52 weeks"},
		{"8760h", "1 year"},
		{"17519h", "1 year"},
		{"17520h", "2 years"},
		{"26279h", "2 years"},
		{"26280h", "3 years"},
		{"201479h", "22 years"},
		{"201480h", "23 years"},
		{"1m0s", "1 minute"},
		{"1m2s", "1 minute"},
		{"3h4m5s", "3 hours"},
		{"6h7m8s9ms", "6 hours"},
		{"0µs", "0 microseconds"},
		{"0ms", "0 milliseconds"},
		{"0s", "0 seconds"},
		{"0m", "0 minutes"},
		{"0h", "0 hours"},
		{"0m1µs", "1 microsecond"},
		{"0m1ms1µs", "1 millisecond"},
		{"0m1ms", "1 millisecond"},
		{"0m1s", "1 second"},
		{"0m1m", "1 minute"},
		{"0m2ms", "2 milliseconds"},
		{"0m2s", "2 seconds"},
		{"0m2m", "2 minutes"},
		{"0m2m3h", "3 hours"},
		{"0m2m34h", "1 day"},
		{"0m56h7m8ms", "2 days"},
		{"-1µs", "-1 microsecond"},
		{"-1ms", "-1 millisecond"},
		{"-1s", "-1 second"},
		{"-1m", "-1 minute"},
		{"-1h", "-1 hour"},
		{"-2µs", "-2 microseconds"},
		{"-2ms", "-2 milliseconds"},
		{"-2s", "-2 seconds"},
		{"-2m", "-2 minutes"},
		{"-2h", "-2 hours"},
		{"-10h", "-10 hours"},
		{"-24h", "-1 day"},
		{"-48h", "-2 days"},
		{"-120h", "-5 days"},
		{"-168h", "-1 week"},
		{"-672h", "-4 weeks"},
		{"-8760h", "-1 year"},
		{"-1m0s", "-1 minute"},
		{"-0m2s", "-2 seconds"},
		{"-0m2m", "-2 minutes"},
		{"-0m2m3h", "-3 hours"},
		{"-0m2m34h", "-1 day"},
		{"-0µs", "-0 microseconds"},
		{"-0ms", "-0 milliseconds"},
		{"-0s", "-0 seconds"},
		{"-0m", "-0 minutes"},
		{"-0h", "-0 hours"},
	}

	for _, table := range testStrings {
		d, err := ParseStringShort(table.test)
		if err != nil {
			t.Errorf("%q", err)
		}
		result := d.String()
		if result != table.expected {
			t.Errorf("d.String() = %q. got %q, expected %q",
				table.test, result, table.expected)
		}
	}
}

// TestInvalidDuration for invalid inputs.
func TestInvalidDuration(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1", ""},
		{"1d", ""},
		{"1w", ""},
		{"1wk", ""},
		{"1y", ""},
		{"", ""},
		{"m1", ""},
		{"1nmd", ""},
		{"0", ""},
		{"-0", ""},
	}

	for _, table := range testStrings {
		_, err := ParseString(table.test)
		if err == nil {
			t.Errorf("ParseString(%q). got %q, expected %q",
				table.test, err, table.expected)
		}
	}

	for _, table := range testStrings {
		_, err := ParseStringShort(table.test)
		if err == nil {
			t.Errorf("ParseString(%q). got %q, expected %q",
				table.test, err, table.expected)
		}
	}
}

// Benchmarks

func BenchmarkParse(b *testing.B) {
	for n := 1; n < b.N; n++ {
		Parse(time.Duration(n) * time.Hour)
	}
}

func BenchmarkParseStringShort(b *testing.B) {
	for n := 1; n < b.N; n++ {
		_, err := ParseStringShort(fmt.Sprintf("%dh", n))
		if err != nil {
			b.Fatal("Benchmark could not complete.", err)
		}
	}
}

func BenchmarkParseString(b *testing.B) {
	for n := 1; n < b.N; n++ {
		_, err := ParseString(fmt.Sprintf("%dh", n))
		if err != nil {
			b.Fatal("Benchmark could not complete.", err)
		}
	}
}
