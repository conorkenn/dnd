package classes

func initWarlock() Class {
	x := new(Class)
	x.ClassName = WARLOCK
	x.HitDie = "d8"
	x.PrimaryAbility = "charisma"
	x.Saves = []string{"wisdom", "charisma"}
	return *x
}
