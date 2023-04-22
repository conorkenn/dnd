package classes

func initRanger() Class {
	x := new(Class)
	x.ClassName = RANGER
	x.HitDie = "d10"
	x.PrimaryAbility = "dexterity and wisdom"
	x.Saves = []string{"strength", "dexterity"}
	return *x
}
