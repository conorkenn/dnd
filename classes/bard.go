package classes

func initBard() Class {
	x := new(Class)
	x.ClassName = BARD
	x.HitDie = "d8"
	x.PrimaryAbility = "charisma"
	x.Saves = []string{"dexterity", "charisma"}
	return *x
}
