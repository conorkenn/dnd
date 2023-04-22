package character

import (
	"fmt"

	"github.com/conorkenn/dnd_go/classes"
	"github.com/conorkenn/dnd_go/races"
	"github.com/conorkenn/dnd_go/stats"
)

type Character struct {
	name         string
	level        int
	race         races.Race
	racialTraits []string
	stats        stats.Stats
	class        classes.Class
}

func (c *Character) Init(name string, race races.Race, className classes.ClassName) {
	c.name = name
	c.level = 1
	c.race = race
	c.stats = stats.GenerateInitialStats(race)
	c.racialTraits = races.GenerateRacialTraits(race)
	c.class = classes.InitClass(className)
}

func (c *Character) levelUp() {
	c.level += 1
}

func (c *Character) Print() {
	fmt.Printf("Name %v\nLevel %v %v %v\n", c.name, c.level, c.race, c.class.ClassName)
	fmt.Printf("Primary Ability: %v", c.class.PrimaryAbility)
	fmt.Printf("\nStats\nDexterity: %v\nStrength: %v\nConstitution: %v\nIntelligence: %v\nCharisma: %v\nWisdom: %v", c.stats.Dexterity, c.stats.Strength, c.stats.Constitution, c.stats.Intelligence, c.stats.Charisma, c.stats.Wisdom)
	fmt.Printf("\nSaves: %v", c.class.Saves)
	fmt.Printf("\nRacial Traits:\n %v\n", c.racialTraits)

}
