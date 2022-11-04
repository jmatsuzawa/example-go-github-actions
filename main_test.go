package main

import "testing"

func TestHello(t *testing.T) {
	// Setup
	expected := "Hello World"
	// Exercise
	actual := Hello("World")
	// Verify
	if actual != "Hello World" {
		t.Errorf("actual=%q, expected=%q", actual, expected)
	}
}

func TestGoodby(t *testing.T) {
	// Setup
	expected := "Goodby World"
	// Exercise
	actual := Goodby("World")
	// Verify
	if actual != expected {
		t.Errorf("actual=%q, expected=%q", actual, expected)
	}
}
