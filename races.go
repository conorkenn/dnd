package main

type Race string

const (
	DRAGONBORN Race = "dragonborn"
	DWARF      Race = "dwarf"
	ELF        Race = "elf"
	GNOME      Race = "gnome"
	HALF_ELF   Race = "half-elf"
	HALFLING   Race = "halfling"
	HALF_ORC   Race = "half-orc"
	HUMAN      Race = "human"
	TIEFLING   Race = "tiefling"
)

func generateRacialTraits(r Race) []string {
	var s []string
	switch r {
	case DRAGONBORN:
		s = []string{"Draconic Ancestry", "Breath Weapon", "Damage Resistance"}
	case DWARF:
		s = []string{"Darkvision", "Dwarven Resilience", "Dwarven Combat", "Dwarven Combat Training", "Stonecunning"}
	case ELF:
		s = []string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"}
	case GNOME:
		s = []string{"Darkvision", "Gnome Cunning"}
	case HALF_ELF:
		s = []string{"Darkvision", "Fey Ancestry", "Skill Versatility"}
	case HALFLING:
		s = []string{"Lucky", "Brave", "Hafling Nimbleness"}
	case HALF_ORC:
		s = []string{"Darkvision", "Menacing", "Relentless Endurance", "Savage Attcks"}
	case HUMAN:
		s = []string{"Extra Language"}
	case TIEFLING:
		s = []string{"Darkvision", "Hellish Resistance", "Infernal Legacy"}
	}

	return s
}
