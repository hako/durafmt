package durafmt

import (
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
)

// TestParse for durafmt time.Duration conversion.
func TestParse(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Second, "1 second"},
		{1 * time.Minute, "1 minute"},
		{2 * time.Second, "2 seconds"},
		{2 * time.Minute, "2 minutes"},
		{1 * time.Hour, "1 hour"},
		{2 * time.Hour, "2 hours"},
		{10 * time.Hour, "10 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{120 * time.Hour, "5 days"},
		{168 * time.Hour, "1 week"},
		{672 * time.Hour, "1 month"},
		{8064 * time.Hour, "1 year"},
		{-1 * time.Second, "-1 second"},
		{-10 * time.Second, "-10 seconds"},
	}

	for _, table := range testTimes {
		result := Parse(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q", table.test, result, result, table.expected)
		}
	}
}

// TestParseString for durafmt duration string conversion.
func TestParseString(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1s", "1 second"},
		{"1m", "1 minute"},
		{"2m", "2 minutes"},
		{"1h", "1 hour"},
		{"10h", "10 hours"},
		{"24h", "1 day"},
		{"48h", "2 days"},
		{"120h", "5 days"},
		{"168h", "1 week"},
		{"672h", "1 month"},
		{"8064h", "1 year"},
		{"1m0s", "1 minute"},
		{"1m2s", "1 minute 2 seconds"},
		{"3h4m5s", "3 hours 4 minutes 5 seconds"},
		{"0s", "0 seconds"},
		{"0m", "0 minutes"},
		{"0h", "0 hours"},
		{"0m2s", "2 seconds"},
		{"0m2m", "2 minutes"},
		{"0m2m3h", "3 hours 2 minutes"},
		{"0m2m34h", "1 day 10 hours 2 minutes"},
		{"-1s", "-1 second"},
		{"-1m", "-1 minute"},
		{"-2m", "-2 minutes"},
		{"-1h", "-1 hour"},
		{"-10h", "-10 hours"},
		{"-24h", "-1 day"},
		{"-48h", "-2 days"},
		{"-120h", "-5 days"},
		{"-168h", "-1 week"},
		{"-672h", "-1 month"},
		{"-8064h", "-1 year"},
		{"-1m0s", "-1 minute"},
		{"-0m2s", "-2 seconds"},
		{"-0m2m", "-2 minutes"},
		{"-0m2m3h", "-3 hours 2 minutes"},
		{"-0m2m34h", "-1 day 10 hours 2 minutes"},
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
			t.Errorf("d.String() = %q. got %q, expected %q", table.test, result, table.expected)
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
		{"1mth", ""},
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
			t.Errorf("NewDurable(%q). got %q, expected %q", table.test, err, table.expected)
		}
	}
}
