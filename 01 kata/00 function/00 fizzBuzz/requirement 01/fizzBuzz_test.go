package main

import (
	"testing"
)

func TestGenerateSequenceFrom1To100(t *testing.T) {
	tests := []struct {
		name   string
		index  int
		result int
	}{
		{"First Number", 0, 1},
		{"Number 25", 24, 25},
		{"Number 50", 49, 50},
		{"Number 75", 74, 75},
		{"Last Number", 99, 100},
	}

	got := GenerateSequenceFrom1To100()

	if len(got) != 100 {
		t.Errorf("Length of sequence is %d, want %d", len(got), 100)
	}

	for _, tt := range tests {
		if got[tt.index] != tt.result {
			t.Errorf("%s: %d, want %d", tt.name, got[tt.index], tt.result)
		}
	}
}
