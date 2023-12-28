package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)

	if result != 5 {
		t.Error("O resultado deve ser 5")
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(2, 3)

	if result != -1 {
		t.Error("O resultado deve ser -1")
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 3)

	if result != 6 {
		t.Error("O resultado deve ser 6")
	}
}
