package classes

func initFighter() Class {
	x := new(Class)
	x.ClassName = FIGHTER
	x.HitDie = "d10"
	return *x
}
