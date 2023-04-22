package classes

func initFighter() Class {
	x := new(Class)
	x.ClassName = FIGHTER
	x.HitDie = "d10"
	x.PrimaryAbility = "strength or dexterity"
	x.Saves = []string{"strength", "constitution"}
	return *x
}
