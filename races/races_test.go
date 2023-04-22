package races

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRacialTraits(t *testing.T) {
	dragonbornRacialTraits := GenerateRacialTraits("dragonborn")
	assert.Equal(t, dragonbornRacialTraits,
		[]string{"Draconic Ancestry", "Breath Weapon", "Damage Resistance"},
		"dragonbornRacialTraits should be equal")

	dwarfRacialTraits := GenerateRacialTraits("dwarf")
	assert.Equal(t, dwarfRacialTraits,
		[]string{"Darkvision", "Dwarven Resilience", "Dwarven Combat",
			"Dwarven Combat Training", "Stonecunning"},
		"dwarfRacialTraits should be equal")

	elfRacialTraits := GenerateRacialTraits("elf")
	assert.Equal(t, elfRacialTraits,
		[]string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"},
		"elfRacialTraits should be equal")

	gnomeRacialTraits := GenerateRacialTraits("gnome")
	assert.Equal(t, gnomeRacialTraits,
		[]string{"Darkvision", "Gnome Cunning"},
		"gnomeRacialTraits should be equal")

	halfElfRacialTraits := GenerateRacialTraits("half-elf")
	assert.Equal(t, halfElfRacialTraits,
		[]string{"Darkvision", "Fey Ancestry", "Skill Versatility"},
		"halfElfRacialTraits should be equal")

	halflingRacialTraits := GenerateRacialTraits("halfling")
	assert.Equal(t, halflingRacialTraits,
		[]string{"Lucky", "Brave", "Hafling Nimbleness"},
		"halflingRacialTraits should be equal")

	halfOrcRacialTraits := GenerateRacialTraits("half-orc")
	assert.Equal(t, halfOrcRacialTraits,
		[]string{"Darkvision", "Menacing", "Relentless Endurance", "Savage Attcks"}, "halfOrcRacialTraits should be equal")

	humanRacialTraits := GenerateRacialTraits("human")
	assert.Equal(t, humanRacialTraits,
		[]string{"Extra Language"}, "humanRacialTraits should be equal")

	tieflingRacialTraits := GenerateRacialTraits("tiefling")
	assert.Equal(t, tieflingRacialTraits,
		[]string{"Darkvision", "Hellish Resistance", "Infernal Legacy"},
		"tieflingRacialTraits should be equal")
}
