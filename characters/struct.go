package characters

type Character struct {
	Name         string
	Health       int
	Mana         int
	Xp           int
	Strength     float64
	Intelligence float64
	Agility      float64
	Defense      float64
	MagicDefense float64
	CritChance   float64
	CritDamage   float64
	CommonSense  float64 // the ability to find other ways to do a task/find a rare item/unlock different dialogues
	StatsBuffs   StatsBuffs
	Localization Localization
	Inventory    Inventory
}

type StatsBuffs struct {
	BuffPerCent float64
	BuffFixed   float64
	BuffTime    int // in seconds
}

type Localization struct {
	X            float64
	Y            float64
	Z            float64
	HitBoxHeight float64
	HitboxWidth  float64
}

type Inventory struct {
	Inventory []items.Items
}
