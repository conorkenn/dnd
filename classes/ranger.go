package classes

func initRanger() Class {
	x := new(Class)
	x.ClassName = RANGER
	x.HitDie = "d10"
	return *x
}
