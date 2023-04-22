package classes

func initWarlock() Class {
	x := new(Class)
	x.ClassName = WARLOCK
	x.HitDie = "d8"
	return *x
}
