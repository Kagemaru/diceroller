package main

import (
	"math/rand"
	"testing"
)

func Test_Setup(t *testing.T) {
	rand.Seed(1)
	wrong_result := rand.Int63()

	rand.Seed(1)
	Setup()
	if rand.Int63() == wrong_result {
		t.Errorf("Setup is supposed to randomize the Seed.")
	}

}
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
