package classes

type ClassName string

const (
	BARBARIAN ClassName = "barbarian"
	BARD      ClassName = "bard"
	CLERIC    ClassName = "cleric"
	DRUID     ClassName = "druid"
	FIGHTER   ClassName = "fighter"
	MONK      ClassName = "monk"
	PALADIN   ClassName = "paladin"
	RANGER    ClassName = "ranger"
	ROGUE     ClassName = "rogue"
	SORCERER  ClassName = "sorcerer"
	WARLOCK   ClassName = "warlock"
	WIZARD    ClassName = "wizard"
)

type Class struct {
	ClassName ClassName
	HitDie    string
}

func InitClass(c ClassName) Class {
	var class Class
	switch c {
	case BARBARIAN:
		class = initBarbarian()
	case BARD:
		class = initBard()
	case CLERIC:
		class = initCleric()
	case DRUID:
		class = initDruid()
	case FIGHTER:
		class = initFighter()
	case MONK:
		class = initMonk()
	case PALADIN:
		class = initPaladin()
	case RANGER:
		class = initRanger()
	case ROGUE:
		class = initRogue()
	case SORCERER:
		class = initSorcerer()
	case WARLOCK:
		class = initWarlock()
	case WIZARD:
		class = initWizard()

	}

	return class
}
