package classes

func initDruid() Class {
	x := new(Class)
	x.ClassName = DRUID
	x.HitDie = "d8"
	return *x
}
