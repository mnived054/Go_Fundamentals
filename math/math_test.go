package math

import (
	"testing"
)

func TestAddition(t *testing.T) {
	output := Addition(4, 6)
	expected := 10

	if output != expected {
		t.Errorf("output: %v, Expected: %v", output, expected)
	}
}
