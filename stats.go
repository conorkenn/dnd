package main

type Stats struct {
	dexterity    int
	strength     int
	constitution int
	intelligence int
	charisma     int
	wisdom       int
}

func generateInitialStats(r Race) Stats {
	s := Stats{}
	switch r {
	case DRAGONBORN:
		s.strength = 2
		s.charisma = 1
	case DWARF:
		s.constitution = 2
	case ELF:
		s.dexterity = 2
	case GNOME:
	case HALF_ELF:
		s.charisma = 2
	case HALFLING:
		s.dexterity = 2
	case HALF_ORC:
		s.strength = 2
		s.constitution = 1
	case HUMAN:
		s.strength = 1
		s.dexterity = 1
		s.constitution = 1
		s.intelligence = 1
		s.wisdom = 1
		s.charisma = 1
	case TIEFLING:
		s.intelligence = 1
		s.charisma = 2
	}
	return s
}

func (s *Stats) increaseDexterity(n int) {
	s.dexterity += n
}

func (s *Stats) increaseStrength(n int) {
	s.strength += n
}

func (s *Stats) increaseConstitution(n int) {
	s.constitution += n
}

func (s *Stats) increaseIntelligence(n int) {
	s.intelligence += n
}

func (s *Stats) increaseCharisma(n int) {
	s.charisma += n
}

func (s *Stats) increaseWisdom(n int) {
	s.wisdom += n
}
