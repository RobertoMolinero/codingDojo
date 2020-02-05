package main

import "testing"

func TestGreet(t *testing.T) {
	bob := "Bob"

	var tests = []struct {
		name string
		args string
		want string
	}{
		{"Requirement 1: Single Person", bob, "Hello, Bob."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}