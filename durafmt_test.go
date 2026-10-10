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

func TestParseWithLimitFirstAdjacentN(t *testing.T) {
	testCases := []struct {
		name     string
		duration time.Duration
		limitN   int
		expected string
	}{
		// Happy path: consecutive units
		{
			name:     "consecutive units limited to 2",
			duration: 2*time.Hour + 15*time.Minute,
			limitN:   2,
			expected: "2 hours 15 minutes",
		},
		{
			name:     "consecutive units limited to 1",
			duration: 2*time.Hour + 15*time.Minute,
			limitN:   1,
			expected: "2 hours",
		},
		{
			name:     "consecutive units limited to 3",
			duration: 2*time.Hour + 15*time.Minute + 30*time.Second,
			limitN:   3,
			expected: "2 hours 15 minutes 30 seconds",
		},
		{
			name:     "consecutive units limited to 2 with 3 available",
			duration: 2*time.Hour + 15*time.Minute + 30*time.Second,
			limitN:   2,
			expected: "2 hours 15 minutes",
		},
		{
			name:     "consecutive year and week limited to 2",
			duration: 8760*time.Hour + 2*168*time.Hour,
			limitN:   2,
			expected: "1 year 2 weeks",
		},
		{
			name:     "consecutive week and day limited to 2",
			duration: 168*time.Hour + 24*time.Hour,
			limitN:   2,
			expected: "1 week 1 day",
		},
		{
			name:     "consecutive day and hour limited to 2",
			duration: 24*time.Hour + 2*time.Hour,
			limitN:   2,
			expected: "1 day 2 hours",
		},
		{
			name:     "consecutive minute and second limited to 2",
			duration: 3*time.Minute + 4*time.Second,
			limitN:   2,
			expected: "3 minutes 4 seconds",
		},
		{
			name:     "consecutive second and millisecond limited to 2",
			duration: 5*time.Second + 6*time.Millisecond,
			limitN:   2,
			expected: "5 seconds 6 milliseconds",
		},
		{
			name:     "consecutive millisecond and microsecond limited to 2",
			duration: 7*time.Millisecond + 8*time.Microsecond,
			limitN:   2,
			expected: "7 milliseconds 8 microseconds",
		},

		// Gap path: non-consecutive units
		{
			name:     "gap between hours and seconds (issue 33 example)",
			duration: 2*time.Hour + 58*time.Second,
			limitN:   2,
			expected: "2 hours",
		},
		{
			name:     "gap between year and second",
			duration: 8760*time.Hour + 2*time.Second,
			limitN:   2,
			expected: "1 year",
		},
		{
			name:     "gap between week and minute",
			duration: 168*time.Hour + 5*time.Minute,
			limitN:   2,
			expected: "1 week",
		},
		{
			name:     "gap between day and second",
			duration: 24*time.Hour + 30*time.Second,
			limitN:   2,
			expected: "1 day",
		},
		{
			name:     "gap between hour and microsecond",
			duration: time.Hour + 5*time.Microsecond,
			limitN:   2,
			expected: "1 hour",
		},
		{
			name:     "gap between minute and microsecond",
			duration: time.Minute + 50*time.Microsecond,
			limitN:   2,
			expected: "1 minute",
		},
		{
			name:     "gap between second and microsecond",
			duration: time.Second + 50*time.Microsecond,
			limitN:   2,
			expected: "1 second",
		},
		{
			name:     "consecutive prefix then gap, requesting limit beyond gap",
			duration: 8760*time.Hour + 2*168*time.Hour + 3*24*time.Hour + 5*time.Minute,
			limitN:   4,
			expected: "1 year 2 weeks 3 days",
		},
		{
			name:     "consecutive prefix then gap, requesting limit at prefix length",
			duration: 8760*time.Hour + 2*168*time.Hour + 3*24*time.Hour + 5*time.Minute,
			limitN:   3,
			expected: "1 year 2 weeks 3 days",
		},
		{
			name:     "consecutive prefix then gap, requesting large limit",
			duration: 8760*time.Hour + 2*168*time.Hour + 3*24*time.Hour + 5*time.Minute,
			limitN:   8,
			expected: "1 year 2 weeks 3 days",
		},

		// Single unit
		{
			name:     "single microsecond with limit 1",
			duration: time.Microsecond,
			limitN:   1,
			expected: "1 microsecond",
		},
		{
			name:     "single microsecond with limit 2",
			duration: time.Microsecond,
			limitN:   2,
			expected: "1 microsecond",
		},
		{
			name:     "single hour with limit 1",
			duration: time.Hour,
			limitN:   1,
			expected: "1 hour",
		},
		{
			name:     "single hour with limit 2",
			duration: time.Hour,
			limitN:   2,
			expected: "1 hour",
		},

		// All 8 units populated
		{
			name: "all units populated with limit 1",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   1,
			expected: "1 year",
		},
		{
			name: "all units populated with limit 2",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   2,
			expected: "1 year 1 week",
		},
		{
			name: "all units populated with limit 3",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   3,
			expected: "1 year 1 week 1 day",
		},
		{
			name: "all units populated with limit 4",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   4,
			expected: "1 year 1 week 1 day 1 hour",
		},
		{
			name: "all units populated with limit 5",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   5,
			expected: "1 year 1 week 1 day 1 hour 1 minute",
		},
		{
			name: "all units populated with limit 6",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   6,
			expected: "1 year 1 week 1 day 1 hour 1 minute 1 second",
		},
		{
			name: "all units populated with limit 7",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   7,
			expected: "1 year 1 week 1 day 1 hour 1 minute 1 second 1 millisecond",
		},
		{
			name: "all units populated with limit 8",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   8,
			expected: "1 year 1 week 1 day 1 hour 1 minute 1 second 1 millisecond 1 microsecond",
		},
		{
			name: "all units populated with limit 10",
			duration: 8760*time.Hour + 168*time.Hour + 24*time.Hour + time.Hour +
				time.Minute + time.Second + time.Millisecond + time.Microsecond,
			limitN:   10,
			expected: "1 year 1 week 1 day 1 hour 1 minute 1 second 1 millisecond 1 microsecond",
		},

		// Boundary values of N
		{
			name:     "limit 0 means no limit (gapped duration fully output)",
			duration: 2*time.Hour + 58*time.Second,
			limitN:   0,
			expected: "2 hours 58 seconds",
		},
		{
			name:     "negative limit means no limit",
			duration: 2*time.Hour + 58*time.Second,
			limitN:   -1,
			expected: "2 hours 58 seconds",
		},

		// Negative durations
		{
			name:     "negative duration with gap limited to 2",
			duration: -(2*time.Hour + 58*time.Second),
			limitN:   2,
			expected: "-2 hours",
		},
		{
			name:     "negative duration without gap limited to 2",
			duration: -(2*time.Hour + 15*time.Minute),
			limitN:   2,
			expected: "-2 hours 15 minutes",
		},
		{
			name:     "negative duration with 3 consecutive units limited to 2",
			duration: -(2*time.Hour + 15*time.Minute + 58*time.Second),
			limitN:   2,
			expected: "-2 hours 15 minutes",
		},
		{
			name:     "negative duration with 3 consecutive units limited to 3",
			duration: -(2*time.Hour + 15*time.Minute + 58*time.Second),
			limitN:   3,
			expected: "-2 hours 15 minutes 58 seconds",
		},

		// Zero duration
		{
			name:     "zero duration with limit 2",
			duration: 0,
			limitN:   2,
			expected: "0 seconds",
		},
		{
			name:     "zero duration with limit 1",
			duration: 0,
			limitN:   1,
			expected: "0 seconds",
		},
		{
			name:     "zero duration with limit 0",
			duration: 0,
			limitN:   0,
			expected: "0 seconds",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Parse(tc.duration).LimitFirstAdjacentN(tc.limitN).String()
			if res != tc.expected {
				t.Errorf("LimitFirstAdjacentN(%d) = %q, expected %q", tc.limitN, res, tc.expected)
			}
			resAlias := Parse(tc.duration).LimitFirstAdjacent(tc.limitN).String()
			if resAlias != tc.expected {
				t.Errorf("LimitFirstAdjacent(%d) = %q, expected %q", tc.limitN, resAlias, tc.expected)
			}
		})
	}
}

func TestParseStringWithLimitFirstAdjacentN(t *testing.T) {
	testCases := []struct {
		input    string
		limitN   int
		expected string
	}{
		{"2h58s", 2, "2 hours"},
		{"2h15m", 2, "2 hours 15 minutes"},
		{"2h15m58s", 2, "2 hours 15 minutes"},
		{"2h15m58s", 3, "2 hours 15 minutes 58 seconds"},
		{"0s", 2, "0 seconds"},
		{"-0s", 2, "-0 seconds"},
		{"0m", 2, "0 minutes"},
		{"0m1s", 2, "1 second"},
		{"0m1ms1µs", 1, "1 millisecond"},
		{"0m1ms1µs", 2, "1 millisecond 1 microsecond"},
		{"0m56h7m8ms", 2, "2 days 8 hours"},
		{"0m56h7m8ms", 3, "2 days 8 hours 7 minutes"},
		{"0m56h7m8ms", 4, "2 days 8 hours 7 minutes"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_limit_%d", tc.input, tc.limitN), func(t *testing.T) {
			d, err := ParseString(tc.input)
			if err != nil {
				t.Fatalf("ParseString(%q) failed: %v", tc.input, err)
			}
			res := d.LimitFirstAdjacentN(tc.limitN).String()
			if res != tc.expected {
				t.Errorf("ParseString(%q).LimitFirstAdjacentN(%d) = %q, expected %q",
					tc.input, tc.limitN, res, tc.expected)
			}
		})
	}
}

func TestLimitFirstAdjacentNWithLimitToUnit(t *testing.T) {
	d := 87593183 * time.Second
	// Total: 2 years 40 weeks 3 days 19 hours 26 minutes 23 seconds
	res := Parse(d).LimitToUnit("days").LimitFirstAdjacentN(2).String()
	expected := "1013 days 19 hours"
	if res != expected {
		t.Errorf("got %q, expected %q", res, expected)
	}

	res = Parse(d).LimitToUnit("hours").LimitFirstAdjacentN(2).String()
	expected = "24331 hours 26 minutes"
	if res != expected {
		t.Errorf("got %q, expected %q", res, expected)
	}
}

func TestInternationalStringWithLimitFirstAdjacentN(t *testing.T) {
	testCases := []struct {
		duration time.Duration
		limitN   int
		expected string
	}{
		{2*time.Hour + 58*time.Second, 2, "2 h"},
		{2*time.Hour + 15*time.Minute, 2, "2 h 15 m"},
		{2*time.Hour + 15*time.Minute + 58*time.Second, 2, "2 h 15 m"},
		{2*time.Hour + 15*time.Minute + 58*time.Second, 3, "2 h 15 m 58 s"},
		{-(2*time.Hour + 58*time.Second), 2, "-2 h"},
		{1001 * time.Microsecond, 1, "1 ms"},
		{1001 * time.Microsecond, 2, "1 ms 1 µs"},
		{170 * time.Hour, 2, "1 w"}, // 170h = 1w 0d 2h, so day is 0, stops at 1w!
		{170 * time.Hour, 0, "1 w 2 h"},
	}

	for _, tc := range testCases {
		res := Parse(tc.duration).LimitFirstAdjacentN(tc.limitN).InternationalString()
		if res != tc.expected {
			t.Errorf("Parse(%v).LimitFirstAdjacentN(%d).InternationalString() = %q, expected %q",
				tc.duration, tc.limitN, res, tc.expected)
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

// TestInternationalString for valid output when using international strings
func TestInternationalString(t *testing.T) {
	var testStrings = []struct {
		test     *Durafmt
		expected string
	}{
		{Parse(5 * time.Microsecond), "5 µs"},
		{Parse(1001 * time.Microsecond), "1 ms 1 µs"},
		{Parse(5 * time.Millisecond), "5 ms"},
		{Parse(1001 * time.Millisecond), "1 s 1 ms"},
		{Parse(1 * time.Second), "1 s"},
		{Parse(5 * time.Second), "5 s"},
		{Parse(65 * time.Second), "1 m 5 s"},
		{Parse(1 * time.Minute), "1 m"},
		{Parse(2 * time.Minute), "2 m"},
		{Parse(65 * time.Minute), "1 h 5 m"},
		{Parse(1 * time.Hour), "1 h"},
		{Parse(2 * time.Hour), "2 h"},
		{Parse(24 * time.Hour), "1 d"},
		{Parse(25 * time.Hour), "1 d 1 h"},
		{Parse(48 * time.Hour), "2 d"},
		{Parse(168 * time.Hour), "1 w"},
		{Parse(170 * time.Hour), "1 w 2 h"},
		{Parse(336 * time.Hour), "2 w"},
	}

	for _, table := range testStrings {
		actual := table.test.InternationalString()
		if actual != table.expected {
			t.Fatalf("actual duration %s does not match expected duration %s", actual, table.expected)
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
