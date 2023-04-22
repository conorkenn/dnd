package classes

func initCleric() Class {
	x := new(Class)
	x.ClassName = CLERIC
	x.HitDie = "d8"
	x.PrimaryAbility = "wisdom"
	x.Saves = []string{"wisdom", "charisma"}
	return *x
}
