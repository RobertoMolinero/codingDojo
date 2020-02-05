package main

import "testing"

func TestGreet(t *testing.T) {
	bob := "Bob"
	jerry := "JERRY"

	var tests = []struct {
		name string
		args *string
		want string
	}{
		{"Requirement 1: Single Person", &bob, "Hello, Bob."},
		{"Requirement 2: Standard", nil, "Hello, my friend."},
		{"Requirement 3: Shouting", &jerry, "HELLO JERRY!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
