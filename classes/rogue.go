package classes

func initRogue() Class {
	x := new(Class)
	x.ClassName = ROGUE
	x.HitDie = "d8"
	x.PrimaryAbility = "dexterity"
	x.Saves = []string{"dexterity", "intelligence"}
	return *x
}
