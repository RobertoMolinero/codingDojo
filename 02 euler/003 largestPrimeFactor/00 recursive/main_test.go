package main

import "testing"

func TestGetLargestPrimeFactor(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"24", args{24}, 3},
		{"13195", args{13195}, 29},
		{"600851475143", args{600851475143}, 6857},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLargestPrimeFactor(tt.args.number); got != tt.want {
				t.Errorf("GetLargestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
