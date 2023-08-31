package main

import (
	"testing"
)

var tests = []struct {
	Name          string
	Dividend      float64
	Divisor       float64
	ExpectedValue float64
	ExpectedErr   bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		result, err := divide(tt.Dividend, tt.Divisor)

		if tt.ExpectedErr {
			if err == nil {
				t.Error("Expected error but didnt get one.")
			}
		} else {
			if err != nil {
				t.Error("Did get an error when wasnt supposed to get one", err.Error())
			}
		}

		if result != tt.ExpectedValue {
			t.Errorf("Expeted value: %v, got value: %v", tt.ExpectedValue, result)
		}
	}
}
