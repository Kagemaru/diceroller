package main

import (
	"math/rand"
	"testing"
)

func Test_Roll(t *testing.T) {
	rand.Seed(1)
	die := Die{1, 2, 0}

	for i := 0; i < 100; i++ {
		result := Roll(die)
		if result != 1 && result != 2 {
			t.Errorf("Result incorrect, go: %d, want: %s", result, "1||2")
		}
	}
}
