package main

import "math/rand"

type Die struct {
	Times    int
	Sides    int
	Modifier int
}

func (d Die) Roll() int {
	var sum int
	for i := 0; i < d.Times; i++ {
		sum += rand.Intn(d.Sides) + 1
	}
	return sum + d.Modifier
}
