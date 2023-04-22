package classes

func initCleric() Class {
	x := new(Class)
	x.ClassName = CLERIC
	x.HitDie = "d8"
	return *x
}
