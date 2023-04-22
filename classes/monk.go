package classes

func initMonk() Class {
	x := new(Class)
	x.ClassName = MONK
	x.HitDie = "d8"
	x.PrimaryAbility = "dexterity and wisdom"
	x.Saves = []string{"strength", "dexterity"}
	return *x
}
