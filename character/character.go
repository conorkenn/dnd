package character

import (
	"github.com/conorkenn/dnd_go/races"
	"github.com/conorkenn/dnd_go/stats"
)

type Character struct {
	name         string
	level        int
	race         races.Race
	racialTraits []string
	stats        stats.Stats
}

func (c *Character) Init(name string, race races.Race) {
	c.name = name
	c.level = 1
	c.race = race
	c.stats = stats.GenerateInitialStats(race)
	c.racialTraits = races.GenerateRacialTraits(race)
}

func (c *Character) levelUp() {
	c.level += 1
}
