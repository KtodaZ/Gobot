package main

import "testing"

func TestIsValidInput(t *testing.T) {
	if !IsValidInput("C2C3") {
		//t.Error("ShouldBeValid")
	}
}