package main

import (
	"regexp"
	"strconv"
)

func parse_dicestring(input string) Die {
	regex := regexp.MustCompile(`^(?:(?P<times>\d+)d|d)?(?P<sides>\d+)(?:\+(?P<bonus>\d+)|\-(?P<malus>\d+))?$`)
	result := regex.FindStringSubmatch(input)

	times, err := strconv.Atoi(result[1])
	if err != nil {
		times = 1
	}
	sides, err := strconv.Atoi(result[2])
	if err != nil {
		sides = 20
	}
	bonus, _ := strconv.Atoi(result[3])
	if err != nil {
		bonus = 0
	}
	malus, _ := strconv.Atoi(result[4])
	if err != nil {
		malus = 0
	}

	var modifier int

	if bonus != 0 {
		modifier = bonus
	} else if malus != 0 {
		modifier = malus * -1
	} else {
		modifier = 0
	}

	return Die{times, sides, modifier}
}
