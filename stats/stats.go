package stats

import "github.com/conorkenn/dnd_go/races"

type Stats struct {
	Dexterity    int
	Strength     int
	Constitution int
	Intelligence int
	Charisma     int
	Wisdom       int
}

func GenerateInitialStats(r races.Race) Stats {
	s := Stats{}
	switch r {
	case races.DRAGONBORN:
		s.Strength = 2
		s.Charisma = 1
	case races.DWARF:
		s.Constitution = 2
	case races.ELF:
		s.Dexterity = 2
	case races.GNOME:
	case races.HALF_ELF:
		s.Charisma = 2
	case races.HALFLING:
		s.Dexterity = 2
	case races.HALF_ORC:
		s.Strength = 2
		s.Constitution = 1
	case races.HUMAN:
		s.Strength = 1
		s.Dexterity = 1
		s.Constitution = 1
		s.Intelligence = 1
		s.Wisdom = 1
		s.Charisma = 1
	case races.TIEFLING:
		s.Intelligence = 1
		s.Charisma = 2
	}
	return s
}

func (s *Stats) increaseDexterity(n int) {
	s.Dexterity += n
}

func (s *Stats) increaseStrength(n int) {
	s.Strength += n
}

func (s *Stats) increaseConstitution(n int) {
	s.Constitution += n
}

func (s *Stats) increaseIntelligence(n int) {
	s.Intelligence += n
}

func (s *Stats) increaseCharisma(n int) {
	s.Charisma += n
}

func (s *Stats) increaseWisdom(n int) {
	s.Wisdom += n
}
