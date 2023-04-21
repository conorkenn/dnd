package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	name := "Suel"
	x := new(Character)
	x.init(name, HUMAN)
	assert.Equal(t, x.level, 1, "level for a freshly initialized character should be level 1")
	assert.Equal(t, x.name, name, "names should be equal")
	assert.Equal(t, x.race, HUMAN, "races should be equal")
	assert.Equal(t, x.stats, Stats{1, 1, 1, 1, 1, 1}, "stat slices should be equal")
	assert.Equal(t, x.racialTraits, []string{"Extra Language"}, "slices should be equal")
}

func TestLevelUp(t *testing.T) {
	name := "Shocks"
	x := new(Character)
	x.init(name, DWARF)
	x.levelUp()
	x.levelUp()
	assert.Equal(t, x.level, 3, "levels should be equal")
}
