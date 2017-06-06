package comap

import "testing"

func TestCoMap_Get(t *testing.T) {
	comap := New()

	// Test the nonexistent element.
	val, ok := comap.Get("Aaron")

	if ok {
		t.Errorf("ok should be false when value does not exist.")
	}

	if val != nil {
		t.Errorf("val that do not exist should be empty.")
	}

	// Test the existed element.
	comap.Set("aaron", "Aaron")	// Set aaron and Aaron where key and value are both string.

	result, ok := comap.Get("aaron")

	if !ok {
		t.Errorf("ok should be true when value exist in the comap")
	}

	if result == nil {
		t.Errorf("result should be Aaron, not nil")
	}
}
