package classes

func initRogue() Class {
	x := new(Class)
	x.ClassName = ROGUE
	x.HitDie = "d8"
	return *x
}
