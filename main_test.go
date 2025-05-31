package main

import (
	"testing"
	"time"
)

func TestParseFuncs(t *testing.T) {
	tests := []struct {
		format   string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{"rfc3339", "2024-06-01T12:34:56Z", time.Date(2024, 6, 1, 12, 34, 56, 0, time.UTC), false},
		{"unix", "1748705359", time.Unix(1748705359, 0).UTC(), false},
		{"unixms", "1748705359000", time.Unix(0, 1748705359000*int64(time.Millisecond)).UTC(), false},
		{"unixus", "1748705359000000", time.Unix(0, 1748705359000000*int64(time.Microsecond)).UTC(), false},
		{"unixns", "1748705359000000000", time.Unix(0, 1748705359000000000).UTC(), false},
		{"unix", "notanumber", time.Time{}, true},
	}

	for _, tt := range tests {
		fn, ok := parseFuncs[tt.format]
		if !ok {
			t.Errorf("parseFuncs missing for format %s", tt.format)
			continue
		}
		got, err := fn(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseFuncs[%s](%q) error = %v, wantErr %v", tt.format, tt.input, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && !got.Equal(tt.expected) {
			t.Errorf("parseFuncs[%s](%q) = %v, want %v", tt.format, tt.input, got, tt.expected)
		}
	}
}

func TestFormatFuncs(t *testing.T) {
	tm := time.Date(2024, 6, 1, 12, 34, 56, 0, time.UTC)
	expected := "2024-06-01T12:34:56Z"
	got := formatFuncs["rfc3339"](tm)
	if got != expected {
		t.Errorf("formatFuncs[rfc3339](%v) = %q, want %q", tm, got, expected)
	}
}
