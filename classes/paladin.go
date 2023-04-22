package classes

func initPaladin() Class {
	x := new(Class)
	x.ClassName = PALADIN
	x.HitDie = "d10"
	x.PrimaryAbility = "strength and charisma"
	x.Saves = []string{"wisdom", "charisma"}
	return *x
}
