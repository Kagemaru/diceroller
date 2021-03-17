package main

import (
	"fmt"
	"testing"
)

func Test_parsedicestring(t *testing.T) {

	for times := 1; times <= 10; times++ {
		for sides := 1; sides <= 10; sides++ {
			for mod := -10; mod <= 10; mod++ {
				sign := ""
				if mod >= 0 {
					sign = "+"
				}

				run_test(t, times, sides, sign, mod)
			}
		}
	}
}

func run_test(t *testing.T, times int, sides int, sign string, mod int) {
	input := fmt.Sprintf("%dd%d%s%d", times, sides, sign, mod)
	result := parse_dicestring(input)

	if times != result.Times {
		t.Errorf("Times not equal! Have: %d, want %d", times, result.Times)
	}

	if sides != result.Sides {
		t.Errorf("Sides not equal! Have: %d, want %d", sides, result.Sides)
	}

	if mod != result.Modifier {
		t.Errorf("Modifier not equal! Have: %d, want %d", mod, result.Modifier)
	}

}
