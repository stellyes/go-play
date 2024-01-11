package main

import "testing"

// run
//  -> go test -coverprofile=coverage.out && go tool cover -html=coverage.out
// to see how your file is covered by your tests

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"random-data-1", 50.0, 5.0, 10.0, false},
	{"throws-error-1", 75.0, 3.0, 25.0, true},
	{"random-data-3", 200.0, 20.0, 10.0, false},
	{"throws-error-4", 80.0, 4.0, 20.0, true},
	{"random-data-5", 120.0, 6.0, 20.0, false},
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("We were supposed to get an error!")
			}
		} else {
			if err != nil {
				t.Error("We weren't supposed to get an error here!")
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got error when should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0.0)
// 	if err == nil {
// 		t.Error("Did not get error when should have")
// 	}
// }
