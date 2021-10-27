package main

import "testing"

func TestIsEven(t *testing.T) {

	if isEven(1) != even(1) {
		t.Errorf("error !!")
	}

	if isEven(2) != even(2) {
		t.Errorf("error !!")
	}

	if isEven(999) != even(999) {
		t.Errorf("error !!")
	}

	if isEven(-999) != even(-999) {
		t.Errorf("error !!")
	}

	if isEven(-998) != even(-998) {
		t.Errorf("error !!")
	}
}

func even(i int) bool {
	return i%2 == 0
}
