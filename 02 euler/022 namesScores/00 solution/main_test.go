package main

import (
	"reflect"
	"testing"
)

func TestGetAlphabeticalValueOfName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Alphabetical Value of Colin", "COLIN", 53},
		{"Alphabetical Value of Mary", "MARY", 57},
		{"Alphabetical Value of Patricia", "PATRICIA", 77},
		{"Alphabetical Value of Linda", "LINDA", 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAlphabeticalValueOfName(tt.input); got != tt.want {
				t.Errorf("GetAlphabeticalValueOfName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNameScore(t *testing.T) {
	tests := []struct {
		name  string
		names []string
		want  int
	}{
		{"", []string{""}, 0},
		{"", []string{"", "", ""}, 0},
		{"", []string{"A", "A", "A"}, 6},
		{"", []string{"A", "B", "C"}, 14},
		{"", []string{"D", "E", "F"}, 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameScore(tt.names); got != tt.want {
				t.Errorf("GetNameScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadTestData(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []string
	}{
		{"Test with abc.txt", "testdata/abc.txt", []string{"A", "B", "C"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadTestData(tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadTestData() = %v, want %v", got, tt.want)
			}
		})
	}
}
