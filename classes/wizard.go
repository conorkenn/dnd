package classes

func initWizard() Class {
	x := new(Class)
	x.ClassName = WIZARD
	x.HitDie = "d6"
	x.PrimaryAbility = "intelligence"
	x.Saves = []string{"intelligence", "wisdom"}
	return *x
}
