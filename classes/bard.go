package classes

func initBard() Class {
	x := new(Class)
	x.ClassName = BARD
	x.HitDie = "d8"
	return *x
}
