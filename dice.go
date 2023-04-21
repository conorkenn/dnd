package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func roll(s string) int {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	numDice := 1

	split := strings.Split(s, "d")
	if split[0] != "" {
		numDice = rollHelper(split[0])
	}

	numSides := rollHelper(split[1])
	total := 0
	for i := 0; i < numDice; i++ {
		total += (r.Intn(numSides) + 1)
	}

	return total
}

func rollHelper(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
