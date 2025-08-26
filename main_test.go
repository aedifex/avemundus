package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Goodbye, world!" // <— Intentional mismatch
	if Hello() != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, Hello())
	}
}
