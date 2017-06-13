package comap

import "testing"

func TestCoMap_Get(t *testing.T) {
	comap := New()

	// Test the nonexistent element.
	val := comap.Get("Aaron")

	if val != nil {
		t.Errorf("val that do not exist should be empty.")
	}

	// Test the existed element.
	comap.Set("aaron", "Aaron") // Set aaron and Aaron where key and value are both string.

	result := comap.Get("aaron")

	if result == nil {
		t.Errorf("result should be Aaron, not nil")
	}
}
