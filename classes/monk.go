package classes

func initMonk() Class {
	x := new(Class)
	x.ClassName = MONK
	x.HitDie = "d8"
	return *x
}
