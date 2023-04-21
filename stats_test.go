package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreaseStats(t *testing.T) {
	s := Stats{1, 1, 1, 1, 1, 1}
	s.increaseDexterity(2)
	s.increaseStrength(3)
	s.increaseConstitution(4)
	s.increaseIntelligence(5)
	s.increaseCharisma(6)
	s.increaseWisdom(7)
	assert.Equal(t, s.dexterity, 3, "The numbers should be the same")
	assert.Equal(t, s.strength, 4, "The numbers should be the same")
	assert.Equal(t, s.constitution, 5, "The numbers should the same")
	assert.Equal(t, s.intelligence, 6, "The numbers should the same")
	assert.Equal(t, s.charisma, 7, "The numbers should the same")
	assert.Equal(t, s.wisdom, 8, "The numbers should be the same")

}

func TestGenerateInitialStats(t *testing.T) {
	dragonbornInitStats := generateInitialStats("dragonborn")
	assert.Equal(t, dragonbornInitStats, Stats{0, 2, 0, 0, 1, 0}, "dragonbornInitStats - slices should be the same")

	dwarfInitStats := generateInitialStats("dwarf")
	assert.Equal(t, dwarfInitStats, Stats{0, 0, 2, 0, 0, 0}, "dwarfInitStats - slices should be the same")

	elfInitStats := generateInitialStats("elf")
	assert.Equal(t, elfInitStats, Stats{2, 0, 0, 0, 0, 0}, "elfInitStats - slices should be the same")

	gnomeInitStats := generateInitialStats("gnome")
	assert.Equal(t, gnomeInitStats, Stats{0, 0, 0, 0, 0, 0}, "gnomeInitStats - slices should be the same")

	halfelfInitStats := generateInitialStats("half-elf")
	assert.Equal(t, halfelfInitStats, Stats{0, 0, 0, 0, 2, 0}, "halfelfInitStats - slices should be the same")

	halflingInitStats := generateInitialStats("halfling")
	assert.Equal(t, halflingInitStats, Stats{2, 0, 0, 0, 0, 0}, "halflingInitStats - slices should be the same")

	halforcInitStats := generateInitialStats("half-orc")
	assert.Equal(t, halforcInitStats, Stats{0, 2, 1, 0, 0, 0}, "halforcInitStats - slices should be the same")

	humanInitStats := generateInitialStats("human")
	assert.Equal(t, humanInitStats, Stats{1, 1, 1, 1, 1, 1}, "humanInitStats - slices should be the same")

	tieflingInitStats := generateInitialStats("tiefling")
	assert.Equal(t, tieflingInitStats, Stats{0, 0, 0, 1, 2, 0}, "tieflingInitStats - slices should be the same")

}
