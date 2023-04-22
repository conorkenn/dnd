package classes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitClass(t *testing.T) {
	x := InitClass(BARBARIAN)
	assert.Equal(t, x.ClassName, BARBARIAN, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d12", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "strength", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "constitution"}, "saves should be equal")

	x = InitClass(BARD)
	assert.Equal(t, x.ClassName, BARD, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"dexterity", "charisma"}, "saves should be equal")

	x = InitClass(CLERIC)
	assert.Equal(t, x.ClassName, CLERIC, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"wisdom", "charisma"}, "saves should be equal")

	x = InitClass(DRUID)
	assert.Equal(t, x.ClassName, DRUID, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"intelligence", "wisdom"}, "saves should be equal")

	x = InitClass(FIGHTER)
	assert.Equal(t, x.ClassName, FIGHTER, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d10", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "strength or dexterity", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "constitution"}, "saves should be equal")

	x = InitClass(MONK)
	assert.Equal(t, x.ClassName, MONK, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "dexterity and wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "dexterity"}, "saves should be equal")

	x = InitClass(PALADIN)
	assert.Equal(t, x.ClassName, PALADIN, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d10", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "strength and charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"wisdom", "charisma"}, "saves should be equal")

	x = InitClass(RANGER)
	assert.Equal(t, x.ClassName, RANGER, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d10", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "dexterity and wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "dexterity"}, "saves should be equal")

	x = InitClass(ROGUE)
	assert.Equal(t, x.ClassName, ROGUE, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "dexterity", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"dexterity", "intelligence"}, "saves should be equal")

	x = InitClass(SORCERER)
	assert.Equal(t, x.ClassName, SORCERER, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d6", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"constitution", "charisma"}, "saves should be equal")

	x = InitClass(WARLOCK)
	assert.Equal(t, x.ClassName, WARLOCK, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"wisdom", "charisma"}, "saves should be equal")

	x = InitClass(WIZARD)
	assert.Equal(t, x.ClassName, WIZARD, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d6", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "intelligence", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"intelligence", "wisdom"}, "saves should be equal")
}

func TestInitBarbarian(t *testing.T) {
	x := initBarbarian()
	assert.Equal(t, x.ClassName, BARBARIAN, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d12", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "strength", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "constitution"}, "saves should be equal")
}

func TestInitBard(t *testing.T) {
	x := initBard()
	assert.Equal(t, x.ClassName, BARD, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"dexterity", "charisma"}, "saves should be equal")
}

func TestInitCleric(t *testing.T) {
	x := initCleric()
	assert.Equal(t, x.ClassName, CLERIC, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"wisdom", "charisma"}, "saves should be equal")
}

func TestInitDruid(t *testing.T) {
	x := initDruid()
	assert.Equal(t, x.ClassName, DRUID, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"intelligence", "wisdom"}, "saves should be equal")
}

func TestInitFighter(t *testing.T) {
	x := initFighter()
	assert.Equal(t, x.ClassName, FIGHTER, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d10", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "strength or dexterity", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "constitution"}, "saves should be equal")
}

func TestInitMonk(t *testing.T) {
	x := initMonk()
	assert.Equal(t, x.ClassName, MONK, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "dexterity and wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "dexterity"}, "saves should be equal")
}

func TestInitPaladin(t *testing.T) {
	x := initPaladin()
	assert.Equal(t, x.ClassName, PALADIN, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d10", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "strength and charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"wisdom", "charisma"}, "saves should be equal")
}

func TestInitRanger(t *testing.T) {
	x := initRanger()
	assert.Equal(t, x.ClassName, RANGER, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d10", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "dexterity and wisdom", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"strength", "dexterity"}, "saves should be equal")
}

func TestInitRogue(t *testing.T) {
	x := initRogue()
	assert.Equal(t, x.ClassName, ROGUE, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "dexterity", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"dexterity", "intelligence"}, "saves should be equal")
}

func TestInitSorcerer(t *testing.T) {
	x := initSorcerer()
	assert.Equal(t, x.ClassName, SORCERER, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d6", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"constitution", "charisma"}, "saves should be equal")
}

func TestInitWarlock(t *testing.T) {
	x := initWarlock()
	assert.Equal(t, x.ClassName, WARLOCK, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d8", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "charisma", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"wisdom", "charisma"}, "saves should be equal")
}

func TestInitWizard(t *testing.T) {
	x := initWizard()
	assert.Equal(t, x.ClassName, WIZARD, "Class names should be equal")
	assert.Equal(t, x.HitDie, "d6", "Hit dies should be equal")
	assert.Equal(t, x.PrimaryAbility, "intelligence", "primary abilities should match")
	assert.Equal(t, x.Saves, []string{"intelligence", "wisdom"}, "saves should be equal")
}
