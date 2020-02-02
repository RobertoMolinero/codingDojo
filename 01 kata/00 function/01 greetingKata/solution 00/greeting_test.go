package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	type args struct {
		pointer []*string
	}

	amy := "Amy"
	bob := "Bob"
	brian := "Brian"
	brianUppercase := "BRIAN"
	charlieDianne := "Charlie, Dianne"
	charlieDianneEscape := "\"Charlie, Dianne\""
	charlotte := "Charlotte"
	jane := "Jane"
	jerry := "JERRY"
	jill := "Jill"

	var tc01 []*string
	tc01 = append(tc01, &bob)

	var tc03 []*string
	tc03 = append(tc03, &jerry)

	var tc04 []*string
	tc04 = append(tc04, &jill)
	tc04 = append(tc04, &jane)

	var tc05 []*string
	tc05 = append(tc05, &amy)
	tc05 = append(tc05, &brian)
	tc05 = append(tc05, &charlotte)

	var tc06 []*string
	tc06 = append(tc06, &amy)
	tc06 = append(tc06, &brianUppercase)
	tc06 = append(tc06, &charlotte)

	var tc07 []*string
	tc07 = append(tc07, &bob)
	tc07 = append(tc07, &charlieDianne)

	var tc08 []*string
	tc08 = append(tc08, &bob)
	tc08 = append(tc08, &charlieDianneEscape)

	var tests = []struct {
		name string
		args args
		want string
	}{
		{"Step 1: Single Person", args{pointer: tc01}, "Hello, Bob."},
		{"Step 2: Nil > Standard", args{pointer: nil}, "Hello, my friend."},
		{"Step 3: Shouting", args{pointer: tc03}, "HELLO JERRY!"},
		{"Step 4: Pair", args{pointer: tc04}, "Hello, Jill and Jane."},
		{"Step 5: List", args{pointer: tc05}, "Hello, Amy, Brian and Charlotte."},
		{"Step 6: List", args{pointer: tc06}, "Hello, Amy and Charlotte. AND HELLO BRIAN!"},
		{"Step 7: String with comma", args{pointer: tc07}, "Hello, Bob, Charlie and Dianne."},
		{"Step 8: Escape", args{pointer: tc08}, "Hello, Bob and Charlie, Dianne."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.pointer...); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
