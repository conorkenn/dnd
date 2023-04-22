package classes

func initSorcerer() Class {
	x := new(Class)
	x.ClassName = SORCERER
	x.HitDie = "d6"
	return *x
}
