package main

import (
	"testing"
)

func TestMostRepeatedElement(t *testing.T) {
	inputData := []string{"apple", "pie", "apple", "red", "red", "red"}
	expectedOutput := "red"

	output := mostRepeatedElement(inputData)

	if output != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, output)
	}
}

func TestMostRepeatedElementEmptyInput(t *testing.T) {
	var inputData []string
	expectedOutput := ""

	output := mostRepeatedElement(inputData)

	if output != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, output)
	}
}

func TestMostRepeatedElementEqualCounts(t *testing.T) {
	inputData := []string{"apple", "pie", "apple", "pie"}
	expectedOutput := "apple" // or "pie", since they both have the same count

	output := mostRepeatedElement(inputData)

	if output != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, output)
	}
}
