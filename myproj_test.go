package myproj

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Sum(3, 7)
	if result != 10 {
		t.Fatalf("want %v, but %v:", 10, result)
	}
}
