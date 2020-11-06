package myproj

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(3, 7)
	if result != 10 {
		t.Fatalf("want %v, but %v:", 10, result)
	}
}
