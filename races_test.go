package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRacialTraits(t *testing.T) {
	dragonbornRacialTraits := generateRacialTraits("dragonborn")
	assert.Equal(t, dragonbornRacialTraits,
		[]string{"Draconic Ancestry", "Breath Weapon", "Damage Resistance"},
		"dragonbornRacialTraits should be equal")

	dwarfRacialTraits := generateRacialTraits("dwarf")
	assert.Equal(t, dwarfRacialTraits,
		[]string{"Darkvision", "Dwarven Resilience", "Dwarven Combat",
			"Dwarven Combat Training", "Stonecunning"},
		"dwarfRacialTraits should be equal")

	elfRacialTraits := generateRacialTraits("elf")
	assert.Equal(t, elfRacialTraits,
		[]string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"},
		"elfRacialTraits should be equal")

	gnomeRacialTraits := generateRacialTraits("gnome")
	assert.Equal(t, gnomeRacialTraits,
		[]string{"Darkvision", "Gnome Cunning"},
		"gnomeRacialTraits should be equal")

	halfElfRacialTraits := generateRacialTraits("half-elf")
	assert.Equal(t, halfElfRacialTraits,
		[]string{"Darkvision", "Fey Ancestry", "Skill Versatility"},
		"halfElfRacialTraits should be equal")

	halflingRacialTraits := generateRacialTraits("halfling")
	assert.Equal(t, halflingRacialTraits,
		[]string{"Lucky", "Brave", "Hafling Nimbleness"},
		"halflingRacialTraits should be equal")

	halfOrcRacialTraits := generateRacialTraits("half-orc")
	assert.Equal(t, halfOrcRacialTraits,
		[]string{"Darkvision", "Menacing", "Relentless Endurance", "Savage Attcks"}, "halfOrcRacialTraits should be equal")

	humanRacialTraits := generateRacialTraits("human")
	assert.Equal(t, humanRacialTraits,
		[]string{"Extra Language"}, "humanRacialTraits should be equal")

	tieflingRacialTraits := generateRacialTraits("tiefling")
	assert.Equal(t, tieflingRacialTraits,
		[]string{"Darkvision", "Hellish Resistance", "Infernal Legacy"},
		"tieflingRacialTraits should be equal")
}
