package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	Setup()
	die := parse_dicestring(os.Args[1])
	fmt.Println(Roll(die))
}

func Setup() {
	rand.Seed(time.Now().UnixNano())
}

func Roll(die Die) int {
	result := die.Roll()

	return result
}
