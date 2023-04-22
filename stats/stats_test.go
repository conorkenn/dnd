package stats

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
	assert.Equal(t, s.Dexterity, 3, "The numbers should be the same")
	assert.Equal(t, s.Strength, 4, "The numbers should be the same")
	assert.Equal(t, s.Constitution, 5, "The numbers should the same")
	assert.Equal(t, s.Intelligence, 6, "The numbers should the same")
	assert.Equal(t, s.Charisma, 7, "The numbers should the same")
	assert.Equal(t, s.Wisdom, 8, "The numbers should be the same")

}

func TestGenerateInitialStats(t *testing.T) {
	dragonbornInitStats := GenerateInitialStats("dragonborn")
	assert.Equal(t, dragonbornInitStats, Stats{0, 2, 0, 0, 1, 0}, "dragonbornInitStats - slices should be the same")

	dwarfInitStats := GenerateInitialStats("dwarf")
	assert.Equal(t, dwarfInitStats, Stats{0, 0, 2, 0, 0, 0}, "dwarfInitStats - slices should be the same")

	elfInitStats := GenerateInitialStats("elf")
	assert.Equal(t, elfInitStats, Stats{2, 0, 0, 0, 0, 0}, "elfInitStats - slices should be the same")

	gnomeInitStats := GenerateInitialStats("gnome")
	assert.Equal(t, gnomeInitStats, Stats{0, 0, 0, 0, 0, 0}, "gnomeInitStats - slices should be the same")

	halfelfInitStats := GenerateInitialStats("half-elf")
	assert.Equal(t, halfelfInitStats, Stats{0, 0, 0, 0, 2, 0}, "halfelfInitStats - slices should be the same")

	halflingInitStats := GenerateInitialStats("halfling")
	assert.Equal(t, halflingInitStats, Stats{2, 0, 0, 0, 0, 0}, "halflingInitStats - slices should be the same")

	halforcInitStats := GenerateInitialStats("half-orc")
	assert.Equal(t, halforcInitStats, Stats{0, 2, 1, 0, 0, 0}, "halforcInitStats - slices should be the same")

	humanInitStats := GenerateInitialStats("human")
	assert.Equal(t, humanInitStats, Stats{1, 1, 1, 1, 1, 1}, "humanInitStats - slices should be the same")

	tieflingInitStats := GenerateInitialStats("tiefling")
	assert.Equal(t, tieflingInitStats, Stats{0, 0, 0, 1, 2, 0}, "tieflingInitStats - slices should be the same")

}
