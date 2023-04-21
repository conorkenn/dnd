package main

type Character struct {
	name         string
	level        int
	race         Race
	racialTraits []string
	stats        Stats
}

func (c *Character) init(name string, race Race) {
	c.name = name
	c.level = 1
	c.race = race
	c.stats = generateInitialStats(race)
	c.racialTraits = generateRacialTraits(race)
}

func (c *Character) levelUp() {
	c.level += 1
}
