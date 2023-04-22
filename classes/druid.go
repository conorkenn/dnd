package classes

func initDruid() Class {
	x := new(Class)
	x.ClassName = DRUID
	x.HitDie = "d8"
	x.PrimaryAbility = "wisdom"
	x.Saves = []string{"intelligence", "wisdom"}
	return *x
}
