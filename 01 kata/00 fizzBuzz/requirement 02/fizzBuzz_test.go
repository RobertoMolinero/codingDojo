package main

import (
	"reflect"
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

func TestFizz(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []string
	}{
		{"First 3 Numbers", []int{1, 2, 3}, []string{"1", "2", "Fizz"}},
		{"Middle of the sequence", []int{48, 49, 50, 51, 52}, []string{"Fizz", "49", "50", "Fizz", "52"}},
		{"Last 3 Numbers", []int{98, 99, 100}, []string{"98", "Fizz", "100"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fizz(tt.args); !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("Length of fizz() = %d, want %d", len(got), len(tt.want))
			}
			if got := Fizz(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizz() = %v, want %v", got, tt.want)
			}
		})
	}
}
