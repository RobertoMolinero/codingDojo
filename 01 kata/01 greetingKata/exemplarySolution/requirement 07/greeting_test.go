package main

import "testing"

func TestGreet(t *testing.T) {
	amy := "Amy"
	bob := "Bob"
	brian := "Brian"
	brianUppercase := "BRIAN"
	charlieDianne := "Charlie, Dianne"
	charlotte := "Charlotte"
	jane := "Jane"
	jerry := "JERRY"
	jill := "Jill"

	var tests = []struct {
		name string
		args []*string
		want string
	}{
		{"Requirement 1: Single Person", []*string{&bob}, "Hello, Bob."},
		{"Requirement 2: Standard", nil, "Hello, my friend."},
		{"Requirement 3: Shouting", []*string{&jerry}, "HELLO JERRY!"},
		{"Requirement 4: Pair", []*string{&jill, &jane}, "Hello, Jill and Jane."},
		{"Requirement 5: List", []*string{&amy, &brian, &charlotte}, "Hello, Amy, Brian and Charlotte."},
		{"Requirement 6: List", []*string{&amy, &brianUppercase, &charlotte}, "Hello, Amy and Charlotte. AND HELLO BRIAN!"},
		{"Requirement 7: String with comma", []*string{&bob, &charlieDianne}, "Hello, Bob, Charlie and Dianne."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args...); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
