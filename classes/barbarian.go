package classes

func initBarbarian() Class {
	x := new(Class)
	x.ClassName = BARBARIAN
	x.HitDie = "d12"
	return *x
}
