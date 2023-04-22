package classes

func initPaladin() Class {
	x := new(Class)
	x.ClassName = PALADIN
	x.HitDie = "d10"
	return *x
}
