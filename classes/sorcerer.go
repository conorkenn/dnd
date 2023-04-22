package classes

func initSorcerer() Class {
	x := new(Class)
	x.ClassName = SORCERER
	x.HitDie = "d6"
	x.PrimaryAbility = "charisma"
	x.Saves = []string{"constitution", "charisma"}
	return *x
}
