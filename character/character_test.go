package character

import (
	"testing"

	"github.com/conorkenn/dnd_go/classes"

	"github.com/conorkenn/dnd_go/races"
	"github.com/conorkenn/dnd_go/stats"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	name := "Suel"
	x := new(Character)
	x.Init(name, races.HUMAN, classes.PALADIN)
	assert.Equal(t, x.level, 1, "level for a freshly initialized character should be level 1")
	assert.Equal(t, x.name, name, "names should be equal")
	assert.Equal(t, x.race, races.HUMAN, "races should be equal")
	assert.Equal(t, x.class, classes.Class{"paladin", "d10"}, "classes should be equal")
	assert.Equal(t, x.stats, stats.Stats{1, 1, 1, 1, 1, 1}, "stat slices should be equal")
	assert.Equal(t, x.racialTraits, []string{"Extra Language"}, "slices should be equal")
}

func TestLevelUp(t *testing.T) {
	name := "Shocks"
	x := new(Character)
	x.Init(name, races.DWARF, classes.RANGER)
	x.levelUp()
	x.levelUp()
	assert.Equal(t, x.level, 3, "levels should be equal")
}
