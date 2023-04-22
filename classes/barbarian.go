package classes

func initBarbarian() Class {
	x := new(Class)
	x.ClassName = BARBARIAN
	x.HitDie = "d12"
	x.PrimaryAbility = "strength"
	x.Saves = []string{"strength", "constitution"}
	return *x
}
